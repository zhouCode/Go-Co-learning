import pandas as pd
import os
import re
import shutil

def read_excel_and_create_folders():
    """读取Excel文件并为每个学生创建文件夹"""
    
    # 读取Excel文件
    excel_file = r'c:\ProgrammingProjects\Co-learning-repo\studentInfo.xlsx'
    
    try:
        # 先删除现有的学生文件夹
        delete_existing_student_folders()
        
        # 尝试读取Excel文件的所有工作表
        excel_data = pd.ExcelFile(excel_file)
        print(f"Excel文件包含的工作表: {excel_data.sheet_names}")
        
        # 读取第一个工作表
        df = pd.read_excel(excel_file, sheet_name=0)
        print(f"数据形状: {df.shape}")
        print(f"列名: {df.columns.tolist()}")
        print("\n前几行数据:")
        print(df.head())
        
        # 分析数据结构，寻找学号、班级、姓名等信息
        print("\n数据详细信息:")
        for col in df.columns:
            print(f"列 '{col}': {df[col].dtype}")
            if not df[col].isna().all():
                print(f"  示例值: {df[col].dropna().head(3).tolist()}")
        
        # 创建文件夹
        create_folders_from_data(df)
        
    except Exception as e:
        print(f"读取Excel文件时出错: {e}")

def delete_existing_student_folders():
    """删除现有的学生文件夹"""
    base_dir = r'c:\ProgrammingProjects\Co-learning-repo'
    
    # 获取目录中的所有项目
    items = os.listdir(base_dir)
    deleted_count = 0
    
    for item in items:
        item_path = os.path.join(base_dir, item)
        
        # 检查是否是学生文件夹（包含学号的文件夹）
        if os.path.isdir(item_path) and re.search(r'\d{10}', item):
            try:
                shutil.rmtree(item_path)
                print(f"删除文件夹: {item}")
                deleted_count += 1
            except Exception as e:
                print(f"删除文件夹 {item} 时出错: {e}")
    
    print(f"总共删除了 {deleted_count} 个学生文件夹\n")

def create_folders_from_data(df):
    """根据数据创建文件夹"""
    
    # 基础目录
    base_dir = r'c:\ProgrammingProjects\Co-learning-repo'
    
    # 尝试识别包含学生信息的列
    student_info_columns = []
    
    for col in df.columns:
        # 检查是否包含学号、班级、姓名等信息
        sample_values = df[col].dropna().astype(str).head(10).tolist()
        
        # 检查是否有类似学号的格式 (数字开头)
        has_student_id = any(re.match(r'^\d{10}', str(val)) for val in sample_values)
        
        # 检查是否包含班级信息
        has_class_info = any('班' in str(val) for val in sample_values)
        
        # 检查是否包含中文姓名
        has_chinese_name = any(re.search(r'[\u4e00-\u9fff]', str(val)) for val in sample_values)
        
        if has_student_id or has_class_info or has_chinese_name:
            student_info_columns.append(col)
            print(f"发现可能的学生信息列: {col}")
            print(f"  示例值: {sample_values[:3]}")
    
    # 如果找到了学生信息列，尝试解析并创建文件夹
    if student_info_columns:
        create_folders_from_identified_columns(df, student_info_columns, base_dir)
    else:
        print("未能识别学生信息列，请检查数据格式")

def create_folders_from_identified_columns(df, info_columns, base_dir):
    """从识别的列中创建文件夹"""
    
    created_folders = []
    
    for index, row in df.iterrows():
        try:
            # 尝试从各列中提取信息
            student_info = {}
            
            for col in info_columns:
                value = str(row[col]) if pd.notna(row[col]) else ""
                
                # 尝试提取学号 (10位数字)
                student_id_match = re.search(r'\b(\d{10})\b', value)
                if student_id_match:
                    student_info['student_id'] = student_id_match.group(1)
                
                # 尝试提取班级信息
                class_match = re.search(r'(\d+[^_]*班[^_]*)', value)
                if class_match:
                    student_info['class_name'] = class_match.group(1)
                
                # 尝试提取姓名 (中文字符)
                name_match = re.search(r'([^\d_\s]+[\u4e00-\u9fff][^\d_\s]*)', value)
                if name_match:
                    student_info['name'] = name_match.group(1)
            
            # 如果成功提取到必要信息，创建文件夹（新格式：班级_学号_姓名）
            if 'student_id' in student_info and 'class_name' in student_info and 'name' in student_info:
                folder_name = f"{student_info['class_name']}_{student_info['student_id']}_{student_info['name']}"
                folder_path = os.path.join(base_dir, folder_name)
                
                # 创建文件夹
                if not os.path.exists(folder_path):
                    os.makedirs(folder_path)
                    created_folders.append(folder_name)
                    print(f"创建文件夹: {folder_name}")
                else:
                    print(f"文件夹已存在: {folder_name}")
                
                # 创建note.md文件
                create_note_file(folder_path, student_info)
            
        except Exception as e:
            print(f"处理第{index+1}行数据时出错: {e}")
    
    print(f"\n总共创建了 {len(created_folders)} 个文件夹")
    return created_folders

def create_note_file(folder_path, student_info):
    """为学生文件夹创建note.md文件"""
    note_file_path = os.path.join(folder_path, "note.md")
    
    # 创建note.md文件内容
    note_content = f"""# {student_info['name']} 的学习笔记

## 基本信息
- **学号**: {student_info['student_id']}
- **班级**: {student_info['class_name']}
- **姓名**: {student_info['name']}

## 学习记录

### 课程笔记
- 

### 作业记录
- 

### 项目经验
- 

### 学习心得
- 

---
*创建时间: {pd.Timestamp.now().strftime('%Y-%m-%d %H:%M:%S')}*
"""
    
    try:
        with open(note_file_path, 'w', encoding='utf-8') as f:
            f.write(note_content)
        print(f"  └─ 创建note.md文件")
    except Exception as e:
        print(f"  └─ 创建note.md文件失败: {e}")

if __name__ == "__main__":
    read_excel_and_create_folders()