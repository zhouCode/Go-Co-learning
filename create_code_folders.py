#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
为每个学生文件夹创建code文件夹和Go语言helloworld示例程序
"""

import os
import re
from datetime import datetime

def create_go_helloworld(code_folder_path, student_name):
    """
    在指定的code文件夹中创建Go语言的helloworld示例程序
    
    Args:
        code_folder_path (str): code文件夹路径
        student_name (str): 学生姓名
    """
    go_file_path = os.path.join(code_folder_path, "main.go")
    
    go_content = f"""package main

import "fmt"

func main() {{
    fmt.Println("Hello, World!")
    fmt.Println("欢迎 {student_name} 学习Go语言!")
    fmt.Println("这是你的第一个Go程序")
    
    // 基本变量声明
    var name string = "{student_name}"
    var age int = 20
    
    fmt.Printf("学生姓名: %s\\n", name)
    fmt.Printf("年龄: %d\\n", age)
    
    // 简单的函数调用
    greet(name)
}}

// 问候函数
func greet(name string) {{
    fmt.Printf("你好, %s! 欢迎来到Go语言的世界!\\n", name)
}}
"""
    
    try:
        with open(go_file_path, 'w', encoding='utf-8') as f:
            f.write(go_content)
        print(f"✓ 创建Go程序: {go_file_path}")
    except Exception as e:
        print(f"✗ 创建Go程序失败 {go_file_path}: {e}")

def create_code_folders():
    """
    为所有学生文件夹创建code文件夹和Go语言示例程序
    """
    base_dir = os.getcwd()
    student_folder_pattern = re.compile(r'^23区块链\d+班_\d+_(.+)$')
    
    created_count = 0
    
    # 遍历当前目录下的所有文件夹
    for item in os.listdir(base_dir):
        item_path = os.path.join(base_dir, item)
        
        # 检查是否是学生文件夹
        if os.path.isdir(item_path) and student_folder_pattern.match(item):
            # 提取学生姓名
            match = student_folder_pattern.match(item)
            student_name = match.group(1)
            
            # 创建code文件夹
            code_folder_path = os.path.join(item_path, "code")
            
            try:
                os.makedirs(code_folder_path, exist_ok=True)
                print(f"✓ 创建code文件夹: {code_folder_path}")
                
                # 在code文件夹中创建Go语言helloworld程序
                create_go_helloworld(code_folder_path, student_name)
                
                created_count += 1
                
            except Exception as e:
                print(f"✗ 创建文件夹失败 {code_folder_path}: {e}")
    
    print(f"\n总结:")
    print(f"成功为 {created_count} 个学生文件夹创建了code文件夹和Go程序")
    print(f"创建时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")

if __name__ == "__main__":
    print("开始为学生文件夹创建code文件夹和Go语言示例程序...")
    create_code_folders()
    print("完成!")