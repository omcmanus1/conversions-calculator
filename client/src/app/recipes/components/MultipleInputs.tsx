import SelectSh from "@/components/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import {
  METRIC_VOLUME,
  METRIC_WEIGHT,
  US_VOLUME,
  US_WEIGHT,
} from "@/constants/measures";
import { InputFields, SingleInput } from "@/types/conversionTypes";
import { Dispatch, SetStateAction } from "react";

type MultipleInputProps = {
  inputList: SingleInput[];
  setInputList: Dispatch<SetStateAction<SingleInput[]>>;
};

export default function MultipleInputsComp({
  inputList,
  setInputList,
}: MultipleInputProps) {
  const handleInputChange = (
    val: string | number,
    index: number,
    field: InputFields
  ) => {
    const updatedList = inputList.map((item, i) =>
      i === index ? { ...item, [field]: val } : item
    );
    setInputList(updatedList);
  };

  const decideDropdowns = (
    inp: SingleInput,
    selectType: "input" | "output"
  ) => {
    switch (true) {
      case inp.inputSystem === "usa" && inp.type === "weight":
        return selectType === "input" ? US_WEIGHT : METRIC_WEIGHT;
      case inp.inputSystem === "usa" && inp.type === "volume":
        return selectType === "input" ? US_VOLUME : METRIC_VOLUME;
      case inp.inputSystem === "metric" && inp.type === "weight":
        return selectType === "input" ? METRIC_WEIGHT : US_WEIGHT;
      case inp.inputSystem === "metric" && inp.type === "volume":
        return selectType === "input" ? METRIC_VOLUME : US_VOLUME;
      default:
        return ["Please specify a unit"];
    }
  };
  return (
    <Card>
      <CardHeader>
        <CardTitle>Add Ingredients Below...</CardTitle>
      </CardHeader>
      <CardContent>
        {inputList.map((inp, index) => {
          return (
            <div key={`inputList_${index}`} className="flex mb-2">
              <p className="mr-2 mt-2">Ingredient: </p>
              <Input
                className="mr-2"
                value={inp.ingredient}
                onChange={(e) =>
                  handleInputChange(e.target.value, index, "ingredient")
                }
              />
              <p className="mr-2 mt-2">Amount: </p>
              <Input
                className="mr-2"
                value={inp.amount || ""}
                type="number"
                onChange={(e) =>
                  handleInputChange(e.target.value, index, "amount")
                }
              />
              <p className="mr-2 mt-2">Unit: </p>
              <SelectSh
                handleChange={(e) => handleInputChange(e, index, "inputUnit")}
                placeholder="Choose..."
                selectContent={decideDropdowns(inp, "input")}
              />
              <p className="ml-2 mr-2 mt-2">Output: </p>
              <SelectSh
                handleChange={(e) => handleInputChange(e, index, "outputUnit")}
                placeholder="Choose..."
                selectContent={decideDropdowns(inp, "output")}
              />
            </div>
          );
        })}
      </CardContent>
    </Card>
  );
}
