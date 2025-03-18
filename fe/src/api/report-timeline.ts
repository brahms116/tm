import { TimelineRequest, TimelineResponse } from "@/contracts";
import { post } from "./post";
import { useQuery, UseQueryResult } from "@tanstack/react-query";

const reportTimelineQuery = async (
  req: TimelineRequest
): Promise<TimelineResponse> => {
  return post("/report-timeline", req) as Promise<TimelineResponse>;
};

export const useReportTimelineQuery = (
  req: TimelineRequest
): UseQueryResult<TimelineResponse> => {
  return useQuery({
    queryKey: ["report-timeline", req],
    queryFn: () => reportTimelineQuery(req),
  });
}
