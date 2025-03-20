import { importTransaction } from "@/api";
import {
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useQueryClient } from "@tanstack/react-query";
import { Upload } from "lucide-react";
import { useDropzone } from "react-dropzone";

export const FileUploadDialog: React.FC<{
  onClose: () => void;
}> = ({ onClose }) => {
  const qc = useQueryClient();
  const handleDrop = async (files: File[]) => {
    for (const file of files) {
      await importTransaction(file);
    }
    qc.invalidateQueries({
      queryKey: ["report-timeline"],
    });
    qc.invalidateQueries({
      queryKey: ["report"],
    });
    onClose();
  };
  const { getRootProps, getInputProps } = useDropzone({
    onDrop: handleDrop,
  });
  return (
    <DialogContent className="sm:max-w-[425px]">
      <DialogHeader>
        <DialogTitle>Upload files</DialogTitle>
        <DialogDescription>Upload your csv files.</DialogDescription>
      </DialogHeader>
      <div
        className="w-full h-72 bg-muted rounded-md flex justify-center items-center p-16 text-center"
        {...getRootProps()}
      >
        <input {...getInputProps()} />
        <Upload size={32} />
      </div>
    </DialogContent>
  );
};
