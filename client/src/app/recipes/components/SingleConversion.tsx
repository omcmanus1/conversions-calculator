import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { singleInput, singleOutput } from "@/types/conversionTypes";
import React, { useState } from "react";
import { postRequest } from "@/api/fetchRequests";
import { usePathname } from "next/navigation";
import SingleInput from "./SingleInput";

export default function SingleConversion({
  weightInputs,
  volumeInputs,
  weightOutputs,
  volumeOutputs,
}: {
  weightInputs: Array<string>;
  volumeInputs: Array<string>;
  weightOutputs: Array<string>;
  volumeOutputs: Array<string>;
  list?: boolean;
}) {
  const pathname = usePathname();
  const subPath = !!pathname.match(/\/recipes\/(.+)/)
    ? pathname.match(/\/recipes\/(.+)/)![1]
    : null;

  const [input, setInput] = useState<singleInput>({
    ingredient: "",
    inputSystem: subPath === "convert-usa" ? "usa" : "metric",
    inputUnit: "",
    outputSystem: subPath === "convert-usa" ? "metric" : "usa",
    outputUnit: "",
    type: "",
    amount: 0,
  });
  const [output, setOutput] = useState<singleOutput>({
    ingredient: "",
    unit: "",
    amount: 0,
  });

  const inputComplete = Object.values(input).every((item) => !!item);

  const handleInput = <K extends keyof singleInput>(
    property: K,
    value: singleInput[K]
  ) => {
    setInput({
      ...input,
      [property]: value === "string" ? value.toLowerCase() : value,
    });
  };

  const handleConversion = async () => {
    let data = { ingredient: "", unit: "", amount: 0 };
    switch (input.inputSystem) {
      case "usa":
        input.type === "volume"
          ? (data = await postRequest("volume-us", input))
          : (data = await postRequest("weight-us", input));
        break;
      case "metric":
        input.type === "volume"
          ? (data = await postRequest("volume-metric", input))
          : (data = await postRequest("weight-metric", input));
        break;
    }
    setOutput(data);
  };

  return (
    <div className="text-center">
      <SingleInput
        input={input}
        weightInputs={weightInputs}
        volumeInputs={volumeInputs}
        weightOutputs={weightOutputs}
        volumeOutputs={volumeOutputs}
        handleInput={handleInput}
      />
      <Button
        className={`mt-3 mb-3 ${inputComplete && "hover:bg-lime-100"}`}
        disabled={!inputComplete}
        variant="outline"
        onClick={handleConversion}
      >
        {input.inputSystem === "usa" ? "Freedom" : "Metric"}
        <ChevronDoubleRight className="w-5" />
        {input.inputSystem === "usa" ? "Metric" : "Freedom"}
      </Button>
      {!!output?.amount && (
        <Card>
          <CardHeader>
            <CardTitle>{output?.ingredient}</CardTitle>
            <CardDescription>
              {`${output?.amount} ${output?.unit}`}
            </CardDescription>
          </CardHeader>
        </Card>
      )}
    </div>
  );
}
