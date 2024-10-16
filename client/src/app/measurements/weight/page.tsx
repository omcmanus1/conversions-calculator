"use client";

import { useState } from "react";
import { postRequest } from "@/api/fetchRequests";
import { Button } from "@/components/ui/button";
import { LoadingSpinner } from "@/components/ui/loading-spinner";
import { BodyWeightTypes, validPaths } from "@/types/bodyWeightTypes";
import { toast } from "sonner";
import BodyWeightInputs from "./components/BodyWeightInputs";

export default function ConvertWeight() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [converted, setConverted] = useState<boolean>(false);
  const [endpointPath, setEndpointPath] = useState<validPaths>("bodyweight-metric");
  const [bodyWeights, setBodyWeights] = useState<BodyWeightTypes>({
    totalLbs: 0,
    totalStone: 0,
    stone: 0,
    lbs: 0,
    kilograms: 0,
  });

  const anyWeightsPopulated = Object.values(bodyWeights).some((weight) => !!weight);
  const buttonText = converted ? "RESET" : "CONVERT";

  const resetBodyWeights = () => {
    setBodyWeights({
      totalLbs: 0,
      totalStone: 0,
      stone: 0,
      lbs: 0,
      kilograms: 0,
    });
    setConverted(false);
  };

  const resetOnError = (errorMessage: string) => {
    toast.error(errorMessage);
    resetBodyWeights();
    setIsLoading(false);
  };

  const handleConversion = async (path: string) => {
    setIsLoading(true);
    if (bodyWeights.stone % 1 !== 0) {
      resetOnError("Stone must be a whole number");
      return;
    }
    const data = await postRequest(path, bodyWeights);
    if (data?.error) {
      resetOnError(data?.error);
      return;
    }
    setBodyWeights({ ...data });
    setConverted(true);
    setIsLoading(false);
  };

  return (
    <>
      <BodyWeightInputs
        bodyWeights={bodyWeights}
        setBodyWeights={setBodyWeights}
        setEndpointPath={setEndpointPath}
        converted={converted}
      />
      <Button
        className={`mt-3 mb-3 ${converted && "text-red-500 hover:text-red-500"}`}
        disabled={isLoading || !anyWeightsPopulated}
        variant="outline"
        onClick={
          !converted ? () => handleConversion(endpointPath) : () => resetBodyWeights()
        }
      >
        {isLoading ? <LoadingSpinner /> : buttonText}
      </Button>
    </>
  );
}
