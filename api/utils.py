# utils.py
import logging
import os
import subprocess

def run_shell_command(command):
    """Run a shell command and capture the output."""
    try:
        output = subprocess.check_output(command, shell=True)
        return output.decode('utf-8')
    except subprocess.CalledProcessError as e:
        logging.error(f"Command '{command}' failed with status {e.returncode}")
        raise e

def get_file_hash(file_path):
    """Return the SHA-256 hash of a file."""
    try:
        with open(file_path, 'rb') as f:
            return ' '.join(['sha256', f'{file_path}:{hashlib.sha256(f.read()).hexdigest()}'])
    except FileNotFoundError:
        raise FileNotFoundError(f"File '{file_path}' not found")

def get_git_hash():
    """Return the current Git commit hash."""
    try:
        return run_shell_command('git rev-parse HEAD')
    except subprocess.CalledProcessError:
        raise RuntimeError('Git repository not found')

import hashlib
import time

def get_file_modified_time(file_path):
    """Return the last modified time of a file in seconds."""
    return os.path.getmtime(file_path)

def get_current_time():
    """Return the current time in seconds."""
    return int(time.time())