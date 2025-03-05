"use client";

import Card from "@/components/Card";
import { GlobeIcon } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import UsersTable from "@/components/UsersTable";

// import { VerificationModal } from "@/components/VerificationModal";

export default function UserAccount() {
  return (
    <>
      <div className="flex items-center gap-5 lg:gap-9 overflow-scroll scrollbar-hide mb-10">
        <Card icon={<GlobeIcon color="#6559F4" size={16} />} iconBg="bg-celebut-violet-light" title="Total Verification Request" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-violet" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#6E9425" size={16} />} iconBg="bg-celebut-green-light" title="Total Approved Requests" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-green" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#FF0000F6" size={16} />} iconBg="bg-celebut-red-light" title="Total Failed Requests" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-red" textColor="text-black" spanColor="text-[#999999]" />
      </div>
      <Tabs defaultValue="all" className="w-auto">
        <TabsList>
          <TabsTrigger value="all">All</TabsTrigger>
          <TabsTrigger value="personal">Personal Account</TabsTrigger>
          <TabsTrigger value="business">Business Account</TabsTrigger>
        </TabsList>
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
    </>
  );
}
