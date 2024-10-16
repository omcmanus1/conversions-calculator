import { Dispatch, SetStateAction } from "react";
import { Input } from "@/components/ui/input";
import { BodyWeightTypes, validPaths } from "@/types/bodyWeightTypes";

type Props = {
  bodyWeights: BodyWeightTypes;
  setBodyWeights: Dispatch<SetStateAction<BodyWeightTypes>>;
  setEndpointPath: Dispatch<SetStateAction<validPaths>>;
  converted: boolean;
};

export default function BodyWeightInputs({
  bodyWeights,
  setBodyWeights,
  setEndpointPath,
  converted,
}: Props) {
  type inputField = {
    heading: string;
    accessor: keyof BodyWeightTypes;
    path: validPaths;
  };

  const inputFields: inputField[] = [
    { heading: "Kilograms", accessor: "kilograms", path: "bodyweight-metric" },
    { heading: "Total Lbs", accessor: "totalLbs", path: "bodyweight-lbs" },
    { heading: "Total Stone", accessor: "totalStone", path: "bodyweight-stone" },
  ];

  const renderTotalInputs = () => {
    return inputFields.map((field) => (
      <>
        <h2 className="text-left p-1">{field.heading}</h2>
        <Input
          type="number"
          placeholder="0"
          readOnly={converted}
          onChange={(e) => {
            setBodyWeights({ ...bodyWeights, [field.accessor]: Number(e.target.value) });
            setEndpointPath(field.path);
          }}
          value={!bodyWeights[field.accessor] ? "" : bodyWeights[field.accessor]}
        />
      </>
    ));
  };

  return (
    <div className="flex flex-col gap-1 start-0">
      {renderTotalInputs()}
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
  );
}
