"use client";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import FilterModal from "@/components/FilterModal";
import BadgeRequestsTable from "@/components/BadgeRequestsTable";
import { LucideListFilter, SearchIcon } from "lucide-react";

export default function UserAccount() {
  const [filterOpen, setFilterOpen] = useState(false);

  const handleApplyFilters = (filters: { status: string[] }) => {
    console.log("Applied filters:", filters);
  };

  return (
    <>
      <div className="flex flex-col lg:flex-row items-center justify-between md:justify-end">
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
      <BadgeRequestsTable />

      <FilterModal open={filterOpen} onOpenChange={setFilterOpen} onApplyFilters={handleApplyFilters} />
    </>
  );
}
