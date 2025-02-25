import { JSX } from "react";

export interface Menu {
  icon: JSX.Element;
  title: string;
  url: string;
}

export interface CardProps {
  title: string;
  percentage: number;
  value: number;
  isPositive: boolean;
  bgColor: string;
  textColor: string;
  spanColor: string;
}
