import customtkinter as ctk
from tkinter import filedialog
import subprocess

ctk.set_appearance_mode("Dark")
ctk.set_default_color_theme("blue")

root = ctk.CTk()
root.title("Forgeware - Dev Toolbox")
root.geometry("600x450")

output_box = ctk.CTkTextbox(root, width=580, height=150)
output_box.pack(pady=10)

def run_tool(tool, file_required=True):
    try:
        args = []
        if tool == "hash":
            file_path = filedialog.askopenfilename(title="Select File")
            if not file_path: return
            args = ["./go_tools/hashcalc/hashcalc", file_path]
        elif tool == "token":
            args = ["./go_tools/gentoken/gentoken"]
        elif tool == "analyze":
            file_path = filedialog.askopenfilename(title="Select File")
            if not file_path: return
            args = ["./go_tools/analyze/analyze", file_path]

        result = subprocess.run(args, capture_output=True, text=True)
        output_box.delete("0.0", "end")
        output_box.insert("0.0", result.stdout)

    except Exception as e:
        output_box.delete("0.0", "end")
        output_box.insert("0.0", f"Error: {e}")

frame = ctk.CTkFrame(root)
frame.pack(pady=10)

ctk.CTkButton(frame, text="Hash File", command=lambda: run_tool("hash")).grid(row=0, column=0, padx=10)
ctk.CTkButton(frame, text="Generate Token", command=lambda: run_tool("token")).grid(row=0, column=1, padx=10)
ctk.CTkButton(frame, text="Analyze File", command=lambda: run_tool("analyze")).grid(row=1, column=0, columnspan=2, pady=10)

root.mainloop()
