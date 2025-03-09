"use client";

import Card from "@/components/Card";
import { GlobeIcon, LucideListFilter, SearchIcon } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import UsersTable from "@/components/UsersTable";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import FilterModal from "@/components/FilterModal";
import { useState } from "react";

// import { VerificationModal } from "@/components/VerificationModal";

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
          <UsersTable account_type="all" />
        </TabsContent>
        <TabsContent value="personal">
          <UsersTable account_type="personal" />
        </TabsContent>
        <TabsContent value="business">
          <UsersTable account_type="business" />
        </TabsContent>
      </Tabs>
      <FilterModal open={filterOpen} onOpenChange={setFilterOpen} onApplyFilters={handleApplyFilters} />
    </>
  );
}
