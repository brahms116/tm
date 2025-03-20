import { getToken } from "./get-token";
import configuration from "@/configuration";

export const importTransaction = async (file: File) => {
  const formData = new FormData();
  formData.append("file", file);
  const res = await fetch(configuration.apiUrl + "/import", {
    headers: {
      Authorization: `${getToken()}`,
    },
    method: "POST",
    body: formData,
  });
  if (!res.ok) {
    throw new Error("Failed to import transactions");
  }
  console.log(await res.json());
};
