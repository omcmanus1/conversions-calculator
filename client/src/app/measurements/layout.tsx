import Navigation from "@/components/navigation";

export default function Layout({ children }: { children: React.ReactNode }) {
  const recipeTabs = [
    { title: "Convert Height", path: "/measurements/height" },
    {
      title: "Convert Weight",
      path: "/measurements/weight",
    },
  ];

  return (
    <div className="flex flex-col items-center text-center">
      <Navigation tabs={recipeTabs}></Navigation>
      <div>{children}</div>
    </div>
  );
}
