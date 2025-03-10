"use client";

import Card from "@/components/Card";
import RevenueChart from "@/components/RevenueChart";
import OSChart from "@/components/OSChart";
import { TrendingUp } from "lucide-react";
import Link from "next/link";
import { pieChartData, lineChartData, barChartData, auditTrail } from "@/lib/constants";
import UsersLoginChart from "@/components/UsersLoginChart";

export default function Dashboard() {
  const revenue = 2000000;
  const percentage = 16;
  return (
    <>
      <div className="flex items-center gap-5 lg:gap-9 overflow-scroll scrollbar-hide mb-10">
        <Card title="Total Users" percentage={16} value={1000000} isPositive={true} bgColor="bg-celebut-green" headingColor="text-white" textColor="text-white" spanColor="text-[#FAFAFA]" />
        <Card title="Total Content" percentage={16} value={1000000} isPositive={false} bgColor="bg-transparent" headingColor="text-black" textColor="text-black" spanColor="text-[#999999]" />
        <Card title="Total Businesses" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-black" textColor="text-black" spanColor="text-[#999999]" />
      </div>
      <div className="grid grid-cols-1 lg:grid-cols-3 min-w-[300px] items-center gap-5 lg:gap-9 mb-10">
        <div className="lg:col-span-2 border border-[#E6E6E6] rounded-[12px] p-4 lg:p-7">
          <div className="flex justify-between items-center mb-2">
            <h4 className="text-black text-base font-semibold">Total Revenue</h4>
            <Link href="/revenue" className="text-celebut-violet text-sm lg:text-base font-semibold">
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

        <div className="border border-[#E6E6E6] rounded-[12px] p-4 lg:p-7 w-full h-full">
          <div className="flex justify-between items-center mb-2">
            <h4 className="text-black text-base font-semibold">Operating Systems</h4>
            <Link href="/operating-system" className="text-celebut-violet text-sm lg:text-base font-semibold">
              View
            </Link>
          </div>

          <OSChart pieChartData={pieChartData} />
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 items-start gap-9">
        <div className="border border-[#E6E6E6] rounded-[12px] px-4 pt-4 lg:px-9 lg:pt-9">
          <div className="flex justify-between items-center mb-[30px]">
            <h4 className="text-black text-lg lg:text-2xl font-semibold">Audit Trail</h4>
            <Link href="/audit-trail" className="text-celebut-violet text-sm lg:text-base font-semibold">
              View details
            </Link>
          </div>
          <ul>
            {auditTrail.map((item) => (
              <li key={item.id} className="flex justify-between mb-6">
                <div>
                  <h5 className="text-base font-semibold mb-2">{item.name}</h5>
                  <p className="text-base font-semibold text-[#8C8CA1]">{item.sector}</p>
                </div>
                <span className="text-celebut-green text-sm font-normal">{item.status}</span>
              </li>
            ))}
          </ul>
        </div>
        <div className="flex flex-col items-center justify-between h-full">
          <div className="border border-[#E6E6E6] rounded-[12px] w-full p-4 lg:p-9 mb-8 lg:mb-0">
            <div className="flex justify-between items-start mb-[30px]">
              <h4 className="text-black text-lg lg:text-2xl font-semibold">
                Total Users/Logins
                <br />
                <span className="text-[#999999] text-sm font-normal">of the week on the Web, Android & IOS</span>
              </h4>
              <Link href="/audit-trail" className="text-celebut-violet text-sm lg:text-base font-semibold">
                View details
              </Link>
            </div>
            <UsersLoginChart barChartData={barChartData} />
          </div>
          <div className="border border-[#E6E6E6] rounded-[12px] w-full p-4 lg:p-9">
            <div className="flex justify-between items-start mb-[30px]">
              <h4 className="text-black text-lg lg:text-2xl font-semibold">
                Total Users/Logins
                <br />
                <span className="text-[#999999] text-sm font-normal">of the week on the Web, Android & IOS</span>
              </h4>
              <Link href="/audit-trail" className="text-celebut-violet text-sm lg:text-base font-semibold">
                View details
              </Link>
            </div>
            <UsersLoginChart barChartData={barChartData} />
          </div>
        </div>
      </div>
    </>
  );
}
