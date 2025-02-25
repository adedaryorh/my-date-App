import { Label, Pie, PieChart, Legend } from "recharts";
import { ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent } from "@/components/ui/chart";
import { useMemo } from "react";
import { ContentType } from "recharts/types/component/DefaultLegendContent";

const chartConfig = {
  users: {
    label: "Users",
  },
  web: {
    label: "Web",
    color: "hsl(var(--chart-7))",
  },
  android: {
    label: "Android",
    color: "hsl(var(--chart-8))",
  },
  IOS: {
    label: "IOS",
    color: "hsl(var(--chart-9))",
  },
} satisfies ChartConfig;

const renderCustomLegend: ContentType = ({ payload }) => {
  return (
    <ul className="flex justify-between items-center gap-5">
      {payload?.map((entry, index: number) => (
        <li key={`item-${index}`} className="flex gap-2 items-center">
          <span className="capitalize text-sm font-semibold">{entry.value}</span>
          <span className="inline-block w-[30px] h-[10px] rounded-sm" style={{ backgroundColor: entry.color }}></span>
        </li>
      ))}
    </ul>
  );
};

export default function OSChart({ pieChartData }: Readonly<{ pieChartData: { os: string; users: number; fill: string }[] }>) {
  const totalUsers = useMemo(() => {
    return pieChartData.reduce((acc, curr) => acc + curr.users, 0);
  }, [pieChartData]);
  return (
    <ChartContainer config={chartConfig} className="aspect-auto h-[300px] w-full">
      <PieChart>
        <ChartTooltip cursor={false} content={<ChartTooltipContent hideLabel />} />
        <Pie data={pieChartData} dataKey="users" nameKey="os" innerRadius={60} strokeWidth={5}>
          <Label
            content={({ viewBox }) => {
              if (viewBox && "cx" in viewBox && "cy" in viewBox) {
                return (
                  <text x={viewBox.cx} y={viewBox.cy} textAnchor="middle" dominantBaseline="middle">
                    <tspan x={viewBox.cx} y={viewBox.cy} className="fill-muted-foreground text-[10px] font-semibold">
                      Total Users
                    </tspan>
                    <tspan x={viewBox.cx} y={(viewBox.cy || 0) + 14} className="fill-foreground text-sm font-bold">
                      {totalUsers.toLocaleString()}
                    </tspan>
                  </text>
                );
              }
            }}
          />
        </Pie>
        <Legend content={renderCustomLegend} className="-translate-y-2 flex-wrap gap-2" />
      </PieChart>
    </ChartContainer>
  );
}
