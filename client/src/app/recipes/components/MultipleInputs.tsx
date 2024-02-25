import XIcon from "@/components/icons/XIcon";
import SelectSh from "@/components/select";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { METRIC_VOLUME, METRIC_WEIGHT, US_VOLUME, US_WEIGHT } from "@/constants/measures";
import { InputFields, InputListProps, RecipeInput } from "@/types/conversionTypes";
import { Fragment } from "react";

export default function MultipleInputsComp({ inputList, setInputList }: InputListProps) {
  const handleInputChange = (val: string | number, index: number, field: InputFields) => {
    if (!(typeof val === "number" && val < 0)) {
      const updatedList = inputList.map((item, i) => {
        return i === index ? { ...item, [field]: val } : item;
      });
      setInputList(updatedList);
    }
  };

  const decideDropdowns = (
    inp: RecipeInput,
    selectType: "input" | "output"
  ): string[] => {
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

  const handleRemoveClick = (index: number) => {
    let newArr = [...inputList];
    newArr.splice(index, 1);
    setInputList(newArr);
  };

  return (
    <Card className="max-h-64 overflow-y-auto">
      <CardHeader>
        <CardTitle>Add Ingredients Below...</CardTitle>
      </CardHeader>
      <CardContent>
        {inputList.map((inp, index) => {
          return (
            <Fragment key={`inputList_${index}`}>
              {index > 0 && <hr className="flex-grow mt-3 md:mt-0 mb-4 md:mb-3" />}
              <div className="flex md:mb-2 flex-col md:flex-row md:items-center">
                <div className="flex flex-row align-center justify-end mb-1">
                  <p className="md:mr-2 mt-2 mr-4">Ingredient: </p>
                  <Input
                    className="md:mr-2 md:w-28 w-44"
                    value={inp.ingredient}
                    onChange={(e) =>
                      handleInputChange(e.target.value, index, "ingredient")
                    }
                  />
                </div>
                <div className="flex flex-row align-center justify-end mb-1 mt-1">
                  <p className="md:mr-2 mt-2 mr-4">Unit: </p>
                  <SelectSh
                    handleChange={(e) => handleInputChange(e, index, "inputUnit")}
                    selectContent={decideDropdowns(inp, "input")}
                    classNames="md:w-28 w-44 mb-0"
                  />
                </div>
                <div className="flex flex-row align-center justify-end mb-1">
                  <p className="md:mr-2 mt-2 mr-4 md:ml-2">Amount: </p>
                  <Input
                    className={`md:w-28 w-44 ${
                      inp.amount < 0 && "border-2 border-red-500"
                    }`}
                    value={inp.amount || ""}
                    type="number"
                    min="0"
                    onChange={(e) =>
                      handleInputChange(Number(e.target.value), index, "amount")
                    }
                  />
                </div>
                <div className="flex flex-row align-center justify-end mb-1 mt-1">
                  <p className="md:mr-2 mt-2 mr-4 md:ml-2">Output: </p>
                  <SelectSh
                    handleChange={(e) => handleInputChange(e, index, "outputUnit")}
                    selectContent={decideDropdowns(inp, "output")}
                    classNames="md:w-28 w-44"
                  />
                </div>
                <Button
                  variant="ghost"
                  className="sm:mt-2 mb-1 md:mt-0 ml-28 md:ml-2 text-center"
                  onClick={() => handleRemoveClick(index)}
                >
                  <span className="mr-1">Remove</span>
                  <XIcon className="h-6 w-6 text-red-500" />
                </Button>
              </div>
            </Fragment>
          );
        })}
      </CardContent>
    </Card>
  );
}
