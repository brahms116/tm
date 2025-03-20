import { getToken } from "./get-token";
import configuration from "@/configuration";

export const post = async (path: string, data: unknown): Promise<unknown> => {
  const response = await fetch(configuration.apiUrl + path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `${getToken()}`,
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    throw new Error(await response.text());
  }

  return response.json();
};
