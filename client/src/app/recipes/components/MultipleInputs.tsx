import SelectSh from "@/components/select";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { METRIC_VOLUME, METRIC_WEIGHT, US_VOLUME, US_WEIGHT } from "@/constants/measures";
import { InputFields, InputListProps, SingleInput } from "@/types/conversionTypes";
import { Fragment } from "react";

export default function MultipleInputsComp({ inputList, setInputList }: InputListProps) {
  const handleInputChange = (val: string | number, index: number, field: InputFields) => {
    const updatedList = inputList.map((item, i) =>
      i === index ? { ...item, [field]: val } : item
    );
    setInputList(updatedList);
  };

  const decideDropdowns = (inp: SingleInput, selectType: "input" | "output") => {
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
    <Card className="max-h-64 overflow-y-auto">
      <CardHeader>
        <CardTitle>Add Ingredients Below...</CardTitle>
      </CardHeader>
      <CardContent>
        {inputList.map((inp, index) => {
          return (
            <Fragment key={`inputList_${index}`}>
              {index > 0 && <hr className="flex-grow mt-3 md:mt-0 mb-4 md:mb-3" />}
              <div className="flex md:mb-2 flex-col md:flex-row">
                <div className="flex flex-row align-center justify-end">
                  <p className="md:mr-2 mt-2 mr-4">Ingredient: </p>
                  <Input
                    className="md:mr-2 md:w-28 w-44"
                    placeholder="..."
                    value={inp.ingredient}
                    onChange={(e) =>
                      handleInputChange(e.target.value, index, "ingredient")
                    }
                  />
                </div>
                <div className="flex flex-row align-center justify-end">
                  <p className="md:mr-2 mt-2 mr-4">Unit: </p>
                  <SelectSh
                    handleChange={(e) => handleInputChange(e, index, "inputUnit")}
                    placeholder="..."
                    selectContent={decideDropdowns(inp, "input")}
                    classNames="md:w-28 w-44 mb-0"
                  />
                </div>
                <div className="flex flex-row align-center justify-end">
                  <p className="md:mr-2 mt-2 mr-4 md:ml-2">Amount: </p>
                  <Input
                    className="md:w-28 w-44"
                    placeholder="..."
                    value={inp.amount || ""}
                    type="number"
                    onChange={(e) =>
                      handleInputChange(Number(e.target.value), index, "amount")
                    }
                  />
                </div>
                <div className="flex flex-row align-center justify-end">
                  <p className="md:mr-2 mt-2 mr-4 md:ml-2">Output: </p>
                  <SelectSh
                    handleChange={(e) => handleInputChange(e, index, "outputUnit")}
                    placeholder="..."
                    selectContent={decideDropdowns(inp, "output")}
                    classNames="md:w-28 w-44"
                  />
                </div>
              </div>
            </Fragment>
          );
        })}
      </CardContent>
    </Card>
  );
}
