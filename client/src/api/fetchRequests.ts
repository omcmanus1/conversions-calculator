import { RecipeInput } from "@/types/conversionTypes";

const devUrl = process.env.NEXT_PUBLIC_DEV_API_URL;
const productionUrl = process.env.NEXT_PUBLIC_PROD_API_URL;

const isProduction = process.env.NODE_ENV === "production";
const apiUrl = isProduction ? productionUrl : devUrl;

export const getRequest = async () => {
  try {
    const res = await fetch(`${apiUrl}/get-encode`);
    if (!res.ok) {
      throw new Error("Failed to fetch conversions");
    }
    const jsonData = await res.json();
    return jsonData[0];
  } catch (err) {
    if (err instanceof Error) {
      return { error: err.message };
    }
  }
};

export const postRequest = async (path: string, data: RecipeInput | RecipeInput[]) => {
  try {
    const res = await fetch(`${apiUrl}/${path}`, {
      method: "POST",
      mode: "cors",
      cache: "default",
      credentials: "same-origin",
      headers: {
        "Content-Type": "application/json",
      },
      redirect: "follow",
      body: JSON.stringify(data),
    });

    if (!res.ok) {
      const errorText = (await res.text()).trim();
      throw new Error(errorText);
    }

    const jsonData = await res.json();
    return jsonData;
  } catch (err) {
    if (err instanceof Error) {
      return { error: err.message };
    }
  }
};
