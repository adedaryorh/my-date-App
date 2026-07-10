import { CartesianGrid, Line, LineChart, XAxis, YAxis } from "recharts";
import { ChartConfig, ChartContainer, ChartTooltip, ChartTooltipContent } from "@/components/ui/chart";

const chartConfig = {
  revenue: {
    label: "Revenue",
    color: "hsl(var(--chart-6))",
  },
} satisfies ChartConfig;

export default function RevenueChart({ lineChartData }: Readonly<{ lineChartData: { day: string; revenue: number }[] }>) {
  return (
    <ChartContainer config={chartConfig} className="aspect-auto h-[250px] w-full">
      <LineChart
        accessibilityLayer
        data={lineChartData}
        margin={{
          left: 12,
          right: 12,
        }}
      >
        <CartesianGrid vertical={false} strokeDasharray="10 10" strokeWidth={2} />
        <XAxis dataKey="day" tickLine={false} axisLine={false} tickMargin={8} minTickGap={32} tickFormatter={(value) => value.slice(0, 3)} />
        <YAxis dataKey="revenue" tickLine={true} axisLine={false} tickMargin={8} tickCount={8} tickFormatter={(value) => value.toLocaleString()} />
        <ChartTooltip content={<ChartTooltipContent className="w-[150px]" nameKey="views" labelFormatter={(value) => value} />} />
        <Line dataKey="revenue" type="monotone" stroke={`var(--color-revenue)`} strokeWidth={2} dot={false} />
      </LineChart>
    </ChartContainer>
  );
}
