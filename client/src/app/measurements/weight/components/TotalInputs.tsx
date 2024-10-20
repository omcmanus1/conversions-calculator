import { Dispatch, SetStateAction } from "react";
import { Input } from "@/components/ui/input";
import { BodyWeightTypes, validPaths } from "@/types/bodyWeightTypes";

type Props = {
  bodyWeights: BodyWeightTypes;
  setBodyWeights: Dispatch<SetStateAction<BodyWeightTypes>>;
  setEndpointPath: Dispatch<SetStateAction<validPaths>>;
  converted: boolean;
};

type inputField = {
  heading: string;
  accessor: keyof BodyWeightTypes;
  path: validPaths;
};

export default function TotalInputs({
  bodyWeights,
  setBodyWeights,
  setEndpointPath,
  converted,
}: Props) {
  const inputFields: inputField[] = [
    { heading: "Kilograms", accessor: "kilograms", path: "bodyweight-metric" },
    { heading: "Total Lbs", accessor: "totalLbs", path: "bodyweight-lbs" },
    { heading: "Total Stone", accessor: "totalStone", path: "bodyweight-stone" },
  ];
  return inputFields.map((field) => (
    <div key={field.path}>
      <h2 className="text-left p-1">{field.heading}</h2>
      <Input
        type="number"
        placeholder="0"
        readOnly={converted}
        onChange={(e) => {
          setBodyWeights({
            ...bodyWeights,
            [field.accessor]: Number(e.target.value),
          });
          setEndpointPath(field.path);
        }}
        value={!bodyWeights[field.accessor] ? "" : bodyWeights[field.accessor]}
      />
    </div>
  ));
}
