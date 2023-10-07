import { singleInput } from "@/types/conversionTypes";

const baseUrl = "http://localhost:8080/api/convert";

export const getRequest = async () => {
  const res = await fetch(`${baseUrl}/get-encode`);
  if (!res.ok) {
    throw new Error("Failed to fetch conversions");
  }
  const jsonData = await res.json();
  return jsonData[0];
};

export const postRequest = async (path: string, data: singleInput) => {
  const res = await fetch(`${baseUrl}/${path}`, {
    method: "POST",
    mode: "cors",
    cache: "no-cache",
    // credentials: "same-origin",
    headers: {
      "Content-Type": "application/json",
    },
    // redirect: "follow",
    // referrerPolicy: "no-referrer",
    body: JSON.stringify(data),
  });
  const jsonData = await res.json();
  return jsonData;
};
