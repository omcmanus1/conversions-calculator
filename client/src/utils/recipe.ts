import { SingleInput } from "@/types/conversionTypes";

export const inputComplete = (input: SingleInput): boolean => {
  if (input.amount < 0) return false;
  return Object.values(input).every((item) => !!item);
};
