"use client";

import { users } from "@/lib/constants";
import { columns } from "@/lib/users/columns";
import { DataTable } from "@/lib/users/data-table";
import { User } from "@/types/interfaces";
import { ColumnDef } from "@tanstack/react-table";
import { useState } from "react";
import { VerificationModal } from "@/components/VerificationModal";
import { BadgeModal } from "@/components/BadgeModal";
import { Eye } from "lucide-react";
import { NotificationModal } from "./NotificationModal";

export default function UsersTable({ account_type }: { account_type: "all" | "personal" | "business" }) {
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isBadgeModalOpen, setIsBadgeModalOpen] = useState(false);
  const [isNotificationModalOpen, setIsNotificationModalOpen] = useState(false);
  const [response, setResponse] = useState<string>("");

  const data = users.filter((user: User) => user.badge_type === account_type || account_type === "all");

  const handleViewClick = (user: User) => {
    setSelectedUser(user);
    setIsModalOpen(true);
  };

  return (
    <div className="container mx-auto py-10">
      <DataTable
        columns={columns.map((column: ColumnDef<User>) => {
          if (column.id === "id") {
            return {
              ...column,
              cell: ({ row }) => (
                <div className="text-right">
                  <button type="button" title="view" onClick={() => handleViewClick(row.original)} className="text-celebut-violet">
                    <Eye color="black" />
                  </button>
                </div>
              ),
            };
          }
          return column;
        })}
        data={data}
      />
      {selectedUser && <VerificationModal isBadgeModalOpen={isBadgeModalOpen} setIsBadgeModalOpen={setIsBadgeModalOpen} data={selectedUser} isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} />}
      {isBadgeModalOpen && <BadgeModal setResponse={setResponse} isNotificationModalOpen={isNotificationModalOpen} setIsNotificationModalOpen={setIsNotificationModalOpen} isOpen={isBadgeModalOpen} onClose={() => setIsBadgeModalOpen(false)} />}
      {isNotificationModalOpen && (
        <NotificationModal
          isOpen={isNotificationModalOpen}
          onClose={() => {
            setIsNotificationModalOpen(false);
          }}
          isSuccess={true}
          response={response}
        />
      )}
    </div>
  );
}
