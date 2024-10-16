"use client";

import { useState } from "react";
import { postRequest } from "@/api/fetchRequests";
import { Input } from "@/components/ui/input";
import { BodyWeightTypes, validPaths } from "@/types/bodyWeightTypes";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { LoadingSpinner } from "@/components/ui/loading-spinner";

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
      <div className="flex flex-col gap-1 start-0">
        <h2 className="text-left p-1">Kilograms</h2>
        <Input
          type="number"
          placeholder="0"
          readOnly={converted}
          onChange={(e) => {
            setBodyWeights({ ...bodyWeights, kilograms: Number(e.target.value) });
            setEndpointPath("bodyweight-metric");
          }}
          value={!bodyWeights.kilograms ? "" : bodyWeights.kilograms}
        />
        <h2 className="text-left p-1">Total Lbs</h2>
        <Input
          type="number"
          placeholder="0"
          readOnly={converted}
          onChange={(e) => {
            setBodyWeights({ ...bodyWeights, totalLbs: Number(e.target.value) });
            setEndpointPath("bodyweight-lbs");
          }}
          value={!bodyWeights.totalLbs ? "" : bodyWeights.totalLbs}
        />
        <h2 className="text-left p-1">Total Stone</h2>
        <Input
          type="number"
          placeholder="0"
          readOnly={converted}
          onChange={(e) => {
            setBodyWeights({ ...bodyWeights, totalStone: Number(e.target.value) });
            setEndpointPath("bodyweight-stone");
          }}
          value={!bodyWeights.totalStone ? "" : bodyWeights.totalStone}
        />
        <div className="flex flex-row gap-1">
          <div className="flex flex-col">
            <h2 className="text-left p-1">Stone</h2>
            <Input
              type="number"
              placeholder="0"
              readOnly={converted}
              onChange={(e) => {
                setBodyWeights({ ...bodyWeights, stone: Number(e.target.value) });
                setEndpointPath("bodyweight-stone");
              }}
              value={!bodyWeights.stone ? "" : bodyWeights.stone}
            />
          </div>
          <div className="flex flex-col">
            <h2 className="text-left p-1">Lbs</h2>
            <Input
              type="number"
              placeholder="0"
              readOnly={converted}
              onChange={(e) => {
                setBodyWeights({ ...bodyWeights, lbs: Number(e.target.value) });
                setEndpointPath("bodyweight-stone");
              }}
              value={!bodyWeights.lbs ? "" : bodyWeights.lbs}
            />
          </div>
        </div>
      </div>
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
