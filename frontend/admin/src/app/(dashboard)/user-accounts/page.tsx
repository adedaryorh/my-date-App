"use client";

import Card from "@/components/Card";
import { GlobeIcon } from "lucide-react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import AccountsTable from "@/components/AccountsTable";

// import { VerificationModal } from "@/components/VerificationModal";

export default function UserAccount() {
  return (
    <>
      <div className="flex items-center gap-5 lg:gap-9 overflow-scroll scrollbar-hide mb-10">
        <Card icon={<GlobeIcon color="#F5BD4B" size={16} />} iconBg="bg-celebut-gold-light" title="Total Accounts" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-gold" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#73C3CB" size={16} />} iconBg="bg-celebut-teal-light" title="Personal Accounts" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-teal" textColor="text-black" spanColor="text-[#999999]" />
        <Card icon={<GlobeIcon color="#6E9425" size={16} />} iconBg="bg-celebut-green-light" title="Business Accounts" percentage={16} value={1000000} isPositive={true} bgColor="bg-transparent" headingColor="text-celebut-green" textColor="text-black" spanColor="text-[#999999]" />
      </div>
      <Tabs defaultValue="all" className="w-auto">
        <TabsList>
          <TabsTrigger value="all">All</TabsTrigger>
          <TabsTrigger value="personal">Personal Account</TabsTrigger>
          <TabsTrigger value="business">Business Account</TabsTrigger>
        </TabsList>
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
    </>
  );
}
