import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function getFileExtension(filename: string) {
    var nameArr = filename.trim().split(".");
    if (nameArr[0] !== "" && nameArr.length === 2)  {
        return nameArr.pop() ?? "";
    }

    return "";
}
