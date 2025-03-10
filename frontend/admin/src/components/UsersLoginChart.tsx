import { Bar, BarChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts";
import { ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent } from "@/components/ui/chart";
import { ContentType } from "recharts/types/component/DefaultLegendContent";
import { BarChartData } from "@/types/interfaces";

const chartConfig = {
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
    <ul className="flex flex-col justify-between items-start gap-5">
      {payload?.map((entry, index: number) => (
        <li key={`item-${index}`} className="flex flex-col gap-2 items-start">
          <span className="capitalize text-sm font-semibold">{entry.value}</span>
          <span className="inline-block w-[30px] h-[10px]" style={{ backgroundColor: entry.color }}></span>
        </li>
      ))}
    </ul>
  );
};

export default function UsersLoginChart({ barChartData }: { barChartData: BarChartData[] }) {
  return (
    <ChartContainer config={chartConfig}>
      <BarChart accessibilityLayer data={barChartData}>
        <CartesianGrid vertical={false} strokeDasharray="10 10" strokeWidth={2} />
        <XAxis dataKey="day" tickLine={false} tickMargin={10} axisLine={false} tickFormatter={(value) => value.slice(0, 3)} />
        <YAxis tickLine={true} axisLine={false} tickMargin={8} tickCount={8} tickFormatter={(value) => value.toLocaleString()} />
        <ChartTooltip cursor={false} content={<ChartTooltipContent indicator="dashed" />} />
        <Bar dataKey="web" fill="var(--color-web)" radius={2} />
        <Bar dataKey="android" fill="var(--color-android)" radius={2} />
        <Bar dataKey="IOS" fill="var(--color-IOS)" radius={2} />
        <Legend content={renderCustomLegend} className="flex-wrap gap-2" layout="vertical" verticalAlign="top" align="right" />
      </BarChart>
    </ChartContainer>
  );
}
