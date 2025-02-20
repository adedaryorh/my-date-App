import 'package:celebut/apps/apps.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

final GlobalKey<NavigatorState> rootNavigation = GlobalKey(debugLabel: 'root');

final goRouterProvider = Provider<GoRouter>((ref) {
  return GoRouter(
    initialLocation: '/',
    navigatorKey: rootNavigation,
    debugLogDiagnostics: true,
    restorationScopeId: 'app',
    redirect: (context, state) {
      return null;
    },
    routes: [
      GoRoute(
        path: '/',
        name: AppRoute.onboarding.name,
        builder: (context, state) => const OnboardingView(),
      ),
      GoRoute(
        path: '/accountType',
        name: AppRoute.accountType.name,
        builder: (context, state) => const AccountTypeView(),
      ),
      GoRoute(
        path: '/signIn',
        name: AppRoute.signIn.name,
        builder: (context, state) => const SignInView(),
      ),
      GoRoute(
        path: '/personalSignUp',
        name: AppRoute.personalSignUp.name,
        builder: (context, state) => const SignUpPersonal(),
      ),
      GoRoute(
        path: '/businessSignUp',
        name: AppRoute.businessSignUp.name,
        builder: (context, state) => const SignUpBusiness(),
      ),
      GoRoute(
        path: '/forgetPassword',
        name: AppRoute.forgetPassword.name,
        builder: (context, state) => const ForgetPasswordView(),
      ),
      GoRoute(
        path: '/verifyPhone',
        name: AppRoute.verifyPhone.name,
        builder: (context, state) => const VerifyPhoneView(),
      ),
      GoRoute(
        path: '/verificationSuccess',
        name: AppRoute.verificationSuccess.name,
        builder: (context, state) => const VerificationSuccessView(),
      ),
    ],
  );
});

enum AppRoute {
  onboarding,
  accountType,
  signIn,
  personalSignUp,
  businessSignUp,
  forgetPassword,
  verifyPhone,
  verificationSuccess
}
