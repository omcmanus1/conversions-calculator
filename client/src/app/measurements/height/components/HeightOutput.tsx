import { Card, CardHeader, CardTitle } from "@/components/ui/card";
import { HeightOutput } from "@/types/heightTypes";

type Props = { isLoading: boolean; output: HeightOutput | {} };

export default function HeightOutputComp({ isLoading, output }: Props) {
  if (isLoading || !Object.keys(output).length) return null;

  const outputText =
    "feet" in output
      ? `${output.feet}ft, ${output.inches}in`
      : "metres" in output
      ? `${output.metres}m`
      : "";

  return (
    <Card>
      <CardHeader>
        <CardTitle className="whitespace-pre-line">{outputText}</CardTitle>
      </CardHeader>
    </Card>
  );
}
