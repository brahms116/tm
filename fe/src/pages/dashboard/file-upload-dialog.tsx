import {
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Upload } from "lucide-react";
import { useDropzone } from "react-dropzone";

export function FileUploadDialog() {
  const handleDrop = (files: File[]) => {};
  const { getRootProps, getInputProps } = useDropzone({
    onDrop: handleDrop,
  });
  return (
    <DialogContent className="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Upload files</DialogTitle>
        <DialogDescription>Upload your csv files.</DialogDescription>
      </DialogHeader>
      <div className="w-full h-72 bg-muted rounded-md flex justify-center items-center p-16 text-center" {...getRootProps()}>
        <input {...getInputProps()} />
        <Upload size={32} />
      </div>
    </DialogContent>
  );
}
