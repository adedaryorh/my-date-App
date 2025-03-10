"use client";

import { accounts } from "@/lib/constants";
import { columns } from "@/lib/accounts/columns";
import { DataTable } from "@/lib/accounts/data-table";
import { Account } from "@/types/interfaces";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";

export default function AccountsTable({ account_type }: { account_type: "all" | "personal" | "business" }) {
  const data = accounts.filter((account: Account) => account.account_type === account_type || account_type === "all");

  return (
    <div className="container mx-auto py-10">
      <DataTable
        columns={columns.map((column: ColumnDef<Account>) => {
          if (column.id === "id") {
            return {
              ...column,
              cell: ({ row }) => (
                <div className="text-left">
                  <Link href="/user-profile" className="text-celebut-violet">
                    View
                  </Link>
                </div>
              ),
            };
          }
          return column;
        })}
        data={data}
      />
    </div>
  );
}
