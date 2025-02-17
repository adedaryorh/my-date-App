import 'package:flutter/material.dart';

class AppTheme {
  ///additional colors
  static const onboard1 = Color(0xffF5BD4B);
  static const onboard2 = Color(0xff73C3CB);
  static const onboard3 = Color(0xff6E9425);
  static ThemeData get light {
    return ThemeData(
      useMaterial3: true,
      fontFamily: 'Open Sans',
      scaffoldBackgroundColor: const Color(0xffFFFFFF),
      colorScheme: lightColorScheme,
      textTheme: textTheme,
      dialogBackgroundColor: const Color(0xffFFFFFF),
      outlinedButtonTheme: outlinedButtonThemeData(),
      elevatedButtonTheme: elevatedButtonThemeData(),
    );
  }
}

const lightColorScheme = ColorScheme(
  brightness: Brightness.light,
  primary: Color(0xffF5BD4B),
  onPrimary: Color(0xff000000),
  secondary: Color(0xffF5BD4B),
  onSecondary: Color(0xff000000),
  error: Color(0xffFF3501),
  onError: Color(0xffFFFFFF),
  surface: Color(0xffFFFFFF),
  onSurface: Color(0xff000000),
);

const TextTheme textTheme = TextTheme(
  headlineLarge: TextStyle(
    fontSize: 28,
    fontWeight: FontWeight.w700,
    color: Color(0xff000000),
  ),
  headlineMedium: TextStyle(
    fontSize: 26,
    fontWeight: FontWeight.w700,
    color: Color(0xff000000),
  ),
  headlineSmall: TextStyle(
    fontSize: 24,
    fontWeight: FontWeight.w700,
    color: Color(0xff000000),
  ),
  titleLarge: TextStyle(
    fontSize: 22,
    fontWeight: FontWeight.w700,
    color: Color(0xff000000),
  ),
  titleMedium: TextStyle(
    fontSize: 20,
    fontWeight: FontWeight.w700,
    color: Color(0xff000000),
  ),
  titleSmall: TextStyle(
    fontSize: 18,
    fontWeight: FontWeight.w700,
    color: Color(0xff000000),
  ),
  bodyLarge: TextStyle(
    fontSize: 16,
    fontWeight: FontWeight.w600,
    color: Color(0xff000000),
  ),
  bodyMedium: TextStyle(
    fontSize: 14,
    fontWeight: FontWeight.w400,
    color: Color(0xff000000),
  ),
  bodySmall: TextStyle(
    fontSize: 12,
    fontWeight: FontWeight.w400,
    color: Color(0xff000000),
  ),
  labelLarge: TextStyle(
    fontSize: 12,
    fontWeight: FontWeight.w600,
    color: Color(0xff000000),
  ),
  labelMedium: TextStyle(
    fontSize: 10,
    fontWeight: FontWeight.w400,
    color: Color(0xff000000),
  ),
  labelSmall: TextStyle(
    fontSize: 10,
    fontWeight: FontWeight.w400,
    color: Color(0xff000000),
  ),
);

OutlinedButtonThemeData outlinedButtonThemeData() => OutlinedButtonThemeData(
      style: ButtonStyle(
        side: WidgetStateProperty.all(
          const BorderSide(
            color: Color(0xffF5BD4B),
          ),
        ),
        minimumSize: WidgetStateProperty.all(const Size(double.maxFinite, 48)),
        shape: WidgetStateProperty.all<OutlinedBorder>(
          RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
        ),
      ),
    );

ElevatedButtonThemeData elevatedButtonThemeData() => ElevatedButtonThemeData(
      style: ButtonStyle(
        backgroundColor: WidgetStateProperty.resolveWith<Color?>(
          (Set<WidgetState> states) {
            if (states.contains(WidgetState.pressed)) {
              return const Color(0xffF5BD4B);
            } else if (states.contains(WidgetState.disabled)) {
              return const Color(0xff878d95);
            }
            return const Color(0xffF5BD4B);
          },
        ),
        minimumSize: WidgetStateProperty.all(const Size(double.maxFinite, 48)),
        shape: WidgetStateProperty.all<OutlinedBorder>(
          RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(10),
          ),
        ),
      ),
    );

InputDecorationTheme inputDecorationTheme() => InputDecorationTheme(
      filled: true,
      fillColor: Colors.white,
      errorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(5),
        borderSide: const BorderSide(color: Colors.red),
      ),
      focusedErrorBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(5),
      ),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(5),
      ),
      disabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(5),
        borderSide: BorderSide(color: Colors.black.withOpacity(0.5)),
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(5),
      ),
    );
