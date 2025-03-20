import { format } from "date-fns";
import { CartesianGrid, XAxis, BarChart, Bar } from "recharts";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

export type TransactionTimelineDataItem = {
  month: Date;
  spending?: number;
  earning?: number;
};

const chartConfig = {
  spending: {
    label: "Spending",
    color: "#ef4444",
  },
  earning: {
    label: "Earning",
    color: "#a3e635",
  },
} satisfies ChartConfig;

export const TransactionTimeline: React.FC<{
  data: TransactionTimelineDataItem[];
}> = ({ data }) => {

  return (
    <ChartContainer className="h-72 w-full" config={chartConfig}>
      <BarChart
        accessibilityLayer
        data={data.map((e) => {
          return {
            month: format(new Date(e.month), "MMM"),
            earning: e.earning,
            spending: e.spending,
          };
        })}
        margin={{
          left: 12,
          right: 12,
          top: 12,
          bottom: 12,
        }}
      >
        <CartesianGrid vertical={false} />
        <XAxis
          dataKey="month"
          tickLine={false}
          axisLine={false}
          tickMargin={8}
          onClick={(e) => {
            console.log(e);
          }}
        />
        <ChartTooltip
          cursor={false}
          content={<ChartTooltipContent hideLabel />}
        />
        <Bar
          dataKey="spending"
          fill="var(--color-spending)"
          radius={4}
          onClick={(e) => {
            console.log(e);
          }}
        />
        <Bar
          dataKey="earning"
          fill="var(--color-earning)"
          radius={4}
          onClick={(e) => {
            console.log(e);
          }}
        />
      </BarChart>
    </ChartContainer>
  );
};
