import { TimelineResponseItem } from "@/contracts";
import { format } from "date-fns";
import { CartesianGrid, XAxis, BarChart, Bar } from "recharts";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

export type TransactionTimelineDataItem ={
  month: Date;
  spendingCents?: number;
  earningCents?: number;
}

const chartConfig = {
  spendingCents: {
    label: "Spending",
    color: "#ef4444",
  },
  earningCents: {
    label: "Earning",
    color: "#a3e635",
  },
} satisfies ChartConfig;

export const Graph: React.FC<{
  data: TimelineResponseItem[];
}> = ({ data }) => {
  return (
    <ChartContainer className="h-72 w-full" config={chartConfig}>
      <BarChart
        accessibilityLayer
        data={data.map((e) => {
          return {
            month: format(new Date(e.month), "MMM"),
            spendingCents: e.summary.spendingCents / 100,
            earningCents: e.summary.earningCents / 100,
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
          onClick={(e) => {console.log(e)}}
        />
        <ChartTooltip
          cursor={false}
          content={<ChartTooltipContent hideLabel />}
        />
        <Bar
          dataKey="spendingCents"
          fill="var(--color-spendingCents)"
          radius={4}
          onClick={(e) => {console.log(e)}}
        />
        <Bar
          dataKey="earningCents"
          fill="var(--color-earningCents)"
          radius={4}
          onClick={(e) => {console.log(e)}}
        />
      </BarChart>
    </ChartContainer>
  );
};
