"use client";

import Card from "@/components/Card";
import RevenueChart from "@/components/RevenueChart";
import OSChart from "@/components/OSChart";
import { TrendingUp } from "lucide-react";
import Link from "next/link";

const pieChartData = [
  { os: "web", users: 150000000, fill: "var(--color-web)" },
  { os: "android", users: 450000000, fill: "var(--color-android)" },
  { os: "IOS", users: 400000000, fill: "var(--color-IOS)" },
];

const lineChartData = [
  { day: "Sunday", revenue: 200000 },
  { day: "Monday", revenue: 305000 },
  { day: "Tuesday", revenue: 457000 },
  { day: "Wednesday", revenue: 730000 },
  { day: "Thursday", revenue: 901200 },
  { day: "Friday", revenue: 1809000 },
  { day: "Saturday", revenue: 2000000 },
];

export default function Dashboard() {
  const revenue = 2000000;
  const percentage = 16;
  return (
    <>
      <div className="flex items-center gap-5 lg:gap-9 overflow-scroll scrollbar-hide mb-10">
        <Card title="Total Users" percentage={16} value={1000000} isPositive={true} bgColor="bg-celebut-green" textColor="text-white" spanColor="text-[#FAFAFA]" />
        <Card title="Total Content" percentage={16} value={1000000} isPositive={false} bgColor="transparent" textColor="text-black" spanColor="text-[##999999]" />
        <Card title="Total Businesses" percentage={16} value={1000000} isPositive={true} bgColor="transparent" textColor="text-black" spanColor="text-[##999999]" />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 min-w-[300px] items-center gap-5 lg:gap-9">
        <div className="lg:col-span-2 border border-[#E6E6E6] rounded-[12px] p-7">
          <div className="flex justify-between items-center mb-2">
            <h4 className="text-black text-base font-semibold">Total Revenue</h4>
            <Link href="/revenue" className="text-celebut-violet text-base font-semibold">
              View details
            </Link>
          </div>
          <div className="flex items-center gap-4">
            <h4 className="font-bold text-[40px] text-black">
              {revenue.toLocaleString("en-US", {
                style: "currency",
                currency: "USD",
                maximumFractionDigits: 0,
              })}
            </h4>

            <div>
              <div className="flex items-center gap-2">
                <TrendingUp color="#9ACD37" size={24} />

                <span className="text-[#9ACD37] text-xs font-semibold">{percentage}%</span>
              </div>
              <div>
                <span className="text-[##999999] font-normal text-[10px]">In last 7 days</span>
              </div>
            </div>
          </div>

          <RevenueChart lineChartData={lineChartData} />
        </div>

        <div className="border border-[#E6E6E6] rounded-[12px] p-7 w-full">
          <div className="flex justify-between items-center mb-2">
            <h4 className="text-black text-base font-semibold">Operating Systems</h4>
            <Link href="/operating-system" className="text-celebut-violet text-base font-semibold">
              View
            </Link>
          </div>

          <OSChart pieChartData={pieChartData} />
        </div>
      </div>
    </>
  );
}
