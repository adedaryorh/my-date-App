"use client";

import Card from "@/components/Card";
import { GlobeIcon, LucideListFilter, SearchIcon } from "lucide-react";
import UsersTable from "@/components/UsersTable";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import FilterModal from "@/components/FilterModal";
import { useState } from "react";
import { cn } from "@/lib/utils";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import Link from "next/link";
import { DropdownTabsProps, TabItem } from "@/types/interfaces";

function DropdownTabs({ tabs, defaultValue, className, selectClassName, contentClassName, setFilterOpen }: DropdownTabsProps) {
  // Use the first tab as default if not specified
  const defaultTab = defaultValue || tabs[0]?.value;
  const [activeTab, setActiveTab] = useState(defaultTab);

  const handleValueChange = (value: string) => {
    setActiveTab(value);
  };

  return (
    <div className={cn("w-full space-y-4", className)}>
      <div className="flex flex-col lg:flex-row items-center justify-between">
        <div className="flex items-center gap-6 flex-1">
          <Select value={activeTab} onValueChange={handleValueChange}>
            <SelectTrigger className={cn("w-full sm:w-[240px] p-5 text-base", selectClassName)}>
              <SelectValue placeholder="Select an account type" />
            </SelectTrigger>
            <SelectContent>
              {tabs.map((tab) => (
                <SelectItem key={tab.value} value={tab.value}>
                  <div className="flex items-center gap-2">{tab.label}</div>
                </SelectItem>
              ))}
            </SelectContent>
          </Select>
          <div className="flex items-center gap-4">
            <div className="relative w-[370px]">
              <Input type="search" className="flex-1 rounded-lg p-0 ps-14 lg:w-[370px] bg-[#ECF1F4]" placeholder="Search" id="search" />
              <Label className="absolute left-3 top-1/2 transform -translate-y-1/2" htmlFor="search">
                <SearchIcon className="text-gray-400" />
              </Label>
            </div>
            <button title="filter" type="button" className="bg-[#ECF1F4] p-3 rounded-lg" onClick={() => setFilterOpen(true)}>
              <LucideListFilter />
            </button>
          </div>
        </div>
        <div className="relative">
          <Link href="/badge-requests" className="bg-black text-white text-center py-3 px-6 rounded-lg font-normal">
            Badge Requests
          </Link>
          <span className="absolute top-0 right-0 transform translate-x-1/2 -translate-y-full bg-red-500 text-white text-xs font-bold rounded-full h-6 w-6 flex items-center justify-center">6</span>
        </div>
      </div>

      <div className={cn("mt-4", contentClassName)}>
        {tabs.map((tab) => (
          <div key={tab.value} className={cn("transition-opacity duration-200", activeTab === tab.value ? "block opacity-100" : "hidden opacity-0")}>
            {tab.content}
          </div>
        ))}
      </div>
    </div>
  );
}

const tabs: TabItem[] = [
  {
    value: "all",
    label: "All Accounts",
    content: <UsersTable account_type="all" />,
  },
  {
    value: "personal",
    label: "Personal Accounts",
    content: <UsersTable account_type="personal" />,
  },
  {
    value: "business",
    label: "Business Accounts",
    content: <UsersTable account_type="business" />,
  },
];

export default function Verification() {
  const [filterOpen, setFilterOpen] = useState(false);

  const handleApplyFilters = (filters: { status: string[] }) => {
    console.log("Applied filters:", filters);
  };
  return (
    <>
      <div className="flex items-center gap-5 lg:gap-9 overflow-scroll scrollbar-hide mb-10">
        <Card icon={<GlobeIcon color="#6559F4" size={16} />} iconBg="bg-celebut-violet-light" title="Total Verification Request" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-violet" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#6E9425" size={16} />} iconBg="bg-celebut-green-light" title="Total Approved Requests" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-green" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#FF0000F6" size={16} />} iconBg="bg-celebut-red-light" title="Total Failed Requests" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-red" textColor="text-black" spanColor="text-[#999999]" />
      </div>
      {/* <Tabs defaultValue="all" className="w-auto">
        <div className="flex flex-col lg:flex-row items-center justify-between">
          <TabsList>
            <TabsTrigger value="all">All</TabsTrigger>
            <TabsTrigger value="personal">Personal Account</TabsTrigger>
            <TabsTrigger value="business">Business Account</TabsTrigger>
          </TabsList>
          <div className="flex items-center md:mt-5 gap-4">
            <div className="relative w-[370px]">
              <Input type="search" className="flex-1 rounded-lg p-2 ps-14 lg:w-[370px] bg-[#ECF1F4]" placeholder="Search" id="search" />
              <Label className="absolute left-3 top-1/2 transform -translate-y-1/2" htmlFor="search">
                <SearchIcon className="text-gray-400" />
              </Label>
            </div>
            <button title="filter" type="button" className="bg-[#ECF1F4] p-3 rounded-lg" onClick={() => setFilterOpen(true)}>
              <LucideListFilter />
            </button>
          </div>
        </div>
        <TabsContent value="all">
          <UsersTable account_type="all" />
        </TabsContent>
        <TabsContent value="personal">
          <UsersTable account_type="personal" />
        </TabsContent>
        <TabsContent value="business">
          <UsersTable account_type="business" />
        </TabsContent>
      </Tabs> */}

      <DropdownTabs tabs={tabs} defaultValue="all" className="my-custom-class" selectClassName="w-[300px]" contentClassName="mt-8" setFilterOpen={setFilterOpen} />

      <FilterModal open={filterOpen} onOpenChange={setFilterOpen} onApplyFilters={handleApplyFilters} />
    </>
  );
}
