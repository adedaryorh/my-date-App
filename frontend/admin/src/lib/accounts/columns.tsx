"use client";

import { Account } from "@/types/interfaces";
import { ColumnDef } from "@tanstack/react-table";
// import { Eye } from "lucide-react";

export const columns: ColumnDef<Account>[] = [
  {
    accessorFn: (row) => row.username,
    id: "username",
    header: "Username",
  },
  {
    accessorFn: (row) => row.email,
    id: "email",
    header: "Email",
  },
  {
    accessorFn: (row) => row.status,
    id: "status",
    header: "Status",
  },
  {
    accessorFn: (row) => row.country,
    id: "country",
    header: "Country",
  },
  {
    accessorFn: (row) => row.id,
    id: "id",
    header: "Action",
  },
];
