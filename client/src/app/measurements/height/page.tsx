"use client";

import { postRequest } from "@/api/fetchRequests";
import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Button } from "@/components/ui/button";
import { LoadingSpinner } from "@/components/ui/loading-spinner";
import { HeightFeet, HeightMetric, HeightOutput, ValidInputs } from "@/types/heightTypes";
import { useState } from "react";
import { toast } from "sonner";
import HeightInputs from "./components/HeightInputs";
import HeightOutputComp from "./components/HeightOutput";

export default function ConvertHeight() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [input, setInput] = useState<ValidInputs | "">("");
  const [heightMetric, setHeightMetric] = useState<HeightMetric>({
    centimetres: 0,
    metres: 0,
  });
  const [heightFeet, setHeightFeet] = useState<HeightFeet>({ feet: 0, inches: 0 });
  const [output, setOutput] = useState<HeightOutput | {}>({});

  const inputComplete =
    input === "feet"
      ? !!heightFeet.feet && !!heightFeet.inches
      : input === "centimetres"
      ? !!heightMetric.centimetres
      : !!heightMetric.metres;

  const handleConversion = async () => {
    setIsLoading(true);
    const data =
      input === "feet"
        ? await postRequest("height-feet", heightFeet)
        : await postRequest("height-metric", heightMetric);

    if (data?.error) {
      setOutput({ feet: 0, inches: 0 });
      toast.error(data?.error);
      setIsLoading(false);
      return;
    }

    setOutput(data);
    setIsLoading(false);
  };

  return (
    <div className="text-center w-[200px]">
      <HeightInputs
        input={input}
        setInput={setInput}
        setHeightFeet={setHeightFeet}
        setHeightMetric={setHeightMetric}
      />
      {input && (
        <Button
          className={`mt-3 mb-3 ${inputComplete && "hover:bg-lime-100"}`}
          disabled={!inputComplete || isLoading}
          variant="outline"
          onClick={handleConversion}
        >
          {isLoading ? (
            <LoadingSpinner />
          ) : (
            <>
              {input === "feet" ? "FEET" : "METRIC"}
              <ChevronDoubleRight className="w-5" />
              {input === "feet" ? "METRIC" : "FEET"}
            </>
          )}
        </Button>
      )}
      <HeightOutputComp isLoading={isLoading} output={output} />
    </div>
  );
}
