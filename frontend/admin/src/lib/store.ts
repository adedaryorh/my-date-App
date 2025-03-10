import { configureStore } from "@reduxjs/toolkit";
import routeReducer from "@/lib/features/routeSlice";

export const store = configureStore({
  reducer: {
    history: routeReducer,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
