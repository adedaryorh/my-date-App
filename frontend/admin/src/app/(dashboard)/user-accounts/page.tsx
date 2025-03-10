"use client";

import Card from "@/components/Card";
import { GlobeIcon, LucideListFilter, SearchIcon } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import AccountsTable from "@/components/AccountsTable";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import FilterModal from "@/components/FilterModal";

export default function UserAccount() {
  const [filterOpen, setFilterOpen] = useState(false);

  const handleApplyFilters = (filters: { status: string[] }) => {
    console.log("Applied filters:", filters);
  };

  return (
    <>
      <div className="flex items-center gap-5 lg:gap-9 overflow-scroll scrollbar-hide mb-10">
        <Card icon={<GlobeIcon color="#F5BD4B" size={16} />} iconBg="bg-celebut-gold-light" title="Total Accounts" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-gold" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#73C3CB" size={16} />} iconBg="bg-celebut-teal-light" title="Personal Accounts" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-teal" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#6E9425" size={16} />} iconBg="bg-celebut-green-light" title="Business Accounts" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-green" textColor="text-black" spanColor="text-[#999999]" />
      </div>
      <Tabs defaultValue="all" className="w-auto">
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
          <AccountsTable account_type="all" />
        </TabsContent>
        <TabsContent value="personal">
          <AccountsTable account_type="personal" />
        </TabsContent>
        <TabsContent value="business">
          <AccountsTable account_type="business" />
        </TabsContent>
      </Tabs>
      <FilterModal open={filterOpen} onOpenChange={setFilterOpen} onApplyFilters={handleApplyFilters} />
    </>
  );
}
