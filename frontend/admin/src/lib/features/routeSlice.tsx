import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface RouteState {
  previousRoute: string;
  currentRoute: string;
}

const initialState: RouteState = {
  previousRoute: "/",
  currentRoute: "/",
};

const RouteSlice = createSlice({
  name: "route",
  initialState,
  reducers: {
    updateRoute(state, action: PayloadAction<string>) {
      return {
        ...state,
        previousRoute: state.currentRoute,
        currentRoute: action.payload,
      };
    },
  },
});

export const { updateRoute } = RouteSlice.actions;
export default RouteSlice.reducer;
