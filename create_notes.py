import os
import pandas as pd
import re

def create_notes_for_existing_folders():
    """为现有的学生文件夹创建note.md文件"""
    
    base_dir = r'c:\ProgrammingProjects\Co-learning-repo'
    
    # 获取所有学生文件夹
    student_folders = []
    for item in os.listdir(base_dir):
        item_path = os.path.join(base_dir, item)
        if os.path.isdir(item_path) and re.match(r'23区块链\d+班_\d{10}_', item):
            student_folders.append(item)
    
    print(f"找到 {len(student_folders)} 个学生文件夹")
    
    created_count = 0
    for folder_name in student_folders:
        folder_path = os.path.join(base_dir, folder_name)
        
        # 解析文件夹名称获取学生信息
        parts = folder_name.split('_')
        if len(parts) >= 3:
            class_name = parts[0]
            student_id = parts[1]
            name = '_'.join(parts[2:])  # 处理姓名中可能包含下划线的情况
            
            student_info = {
                'class_name': class_name,
                'student_id': student_id,
                'name': name
            }
            
            # 创建note.md文件
            note_file_path = os.path.join(folder_path, "note.md")
            
            if not os.path.exists(note_file_path):
                create_note_file(folder_path, student_info)
                created_count += 1
                print(f"为 {folder_name} 创建note.md文件")
            else:
                print(f"{folder_name} 的note.md文件已存在")
    
    print(f"\n总共创建了 {created_count} 个note.md文件")

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
    except Exception as e:
        print(f"  └─ 创建note.md文件失败: {e}")

if __name__ == "__main__":
    create_notes_for_existing_folders()