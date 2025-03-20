import { ReportRequest, ReportResponse } from "@/contracts";

import { post } from "./post";
import { useQuery, UseQueryResult } from "@tanstack/react-query";

export const report = async (req: ReportRequest): Promise<ReportResponse> => {
  return post("/report-period", req) as Promise<ReportResponse>;
};

export const useReportQuery = (
  req: ReportRequest
): UseQueryResult<ReportResponse> => {
  return useQuery({
    queryKey: ["report", req],
    queryFn: () => report(req),
  });
};
