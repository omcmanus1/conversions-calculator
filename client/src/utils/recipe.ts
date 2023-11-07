import { SingleInput } from "@/types/conversionTypes";

export const inputComplete = (input: SingleInput): boolean => {
  return Object.values(input).every((item) => !!item);
};
