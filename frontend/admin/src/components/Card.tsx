import { CardProps } from "@/types/interfaces";
import { TrendingDown, TrendingUp } from "lucide-react";

export default function Card({ title, percentage, value, isPositive, bgColor, textColor, spanColor }: Readonly<CardProps>) {
  return (
    <div className={`p-7 ${bgColor} ${textColor} border border-[#E6E6E6] rounded-[12px] flex-1`}>
      <h3 className="text-base font-semibold mb-7">{title}</h3>
      <div className="flex items-center gap-4">
        <h4 className="text-4xl font-bold">{(+value)?.toLocaleString()}</h4>
        <div>
          <div className="flex items-center justify-between">
            {isPositive ? <TrendingUp color="#9ACD37" size={24} /> : <TrendingDown color="#FF0000F6" size={24} />}
            <span className={`${isPositive ? "text-[#9ACD37]" : "text-celebut-red"} text-xs font-semibold`}>{percentage}%</span>
          </div>
          <div>
            <span className={`${spanColor} font-normal text-[10px]`}>In last 7 days</span>
          </div>
        </div>
      </div>
    </div>
  );
}
