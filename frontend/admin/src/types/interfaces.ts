import { JSX } from "react";

export interface Menu {
  icon: JSX.Element;
  title: string;
  url: string;
}

export interface CardProps {
  title: string;
  icon?: JSX.Element;
  iconBg?: string;
  percentage: number;
  value: number;
  isPositive: boolean;
  bgColor: string;
  headingColor: string;
  textColor: string;
  spanColor: string;
}

export interface PieChartData {
  os: string;
  users: number;
  fill: string;
}

export interface LineChartData {
  day: string;
  revenue: number;
}

export interface BarChartData {
  day: string;
  web: number;
  android: number;
  IOS: number;
}

export interface AuditTrail {
  id: number;
  name: string;
  sector: string;
  status: string;
}

export interface User {
  id: string;
  image: string;
  fullname: string;
  email: string;
  status: "pending" | "approved" | "failed";
  badge_type: string;
  country: string;
  state: string;
  document_type: string;
  application_date: string;
  dob: string;
}

export interface Badge {
  id: number;
  name: string;
  image: JSX.Element;
}

export interface Account {
  id: number;
  username: string;
  email: string;
  status: string;
  country: string;
  account_type: "personal" | "business";
}

export interface FilterModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  onApplyFilters?: (filters: { status: string[] }) => void;
}
