import { postRequest } from "@/api/fetchRequests";
import ChevronDoubleRight from "@/components/icons/ChevronDoubleRight";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  METRIC_VOLUME,
  METRIC_WEIGHT,
  US_VOLUME,
  US_WEIGHT,
} from "@/constants/measures";
import { singleInput, singleOutput } from "@/types/conversionTypes";
import { useState } from "react";
import SingleInput from "./SingleInput";

export default function SingleConversion({
  conversionType,
}: {
  conversionType: "usa" | "metric";
}) {
  const weightInputs = conversionType === "usa" ? US_WEIGHT : METRIC_WEIGHT;
  const weightOutputs = conversionType === "usa" ? METRIC_WEIGHT : US_WEIGHT;
  const volumeInputs = conversionType === "usa" ? US_VOLUME : METRIC_VOLUME;
  const volumeOutputs = conversionType === "usa" ? METRIC_VOLUME : US_VOLUME;

  const [input, setInput] = useState<singleInput>({
    ingredient: "",
    inputSystem: conversionType === "usa" ? "usa" : "metric",
    inputUnit: "",
    outputSystem: conversionType === "usa" ? "metric" : "usa",
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
