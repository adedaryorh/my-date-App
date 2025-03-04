import 'package:celebut/apps/apps.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

final GlobalKey<NavigatorState> rootNavigation = GlobalKey(debugLabel: 'root');
final GlobalKey<NavigatorState> personalShellNavigation =
    GlobalKey(debugLabel: 'personalShell');
final GlobalKey<NavigatorState> businessShellNavigation =
    GlobalKey(debugLabel: 'businessShell');

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
      ShellRoute(
        navigatorKey: personalShellNavigation,
        parentNavigatorKey: rootNavigation,
        restorationScopeId: 'app',
        builder: (context, state, child) =>
            MainEntryPersonalView(key: state.pageKey, child: child),
        routes: [
          GoRoute(
            path: '/pCelebration',
            name: AppRoute.pCelebration.name,
            pageBuilder: (context, state) => MaterialPage(
              child: CelebrationPersonal(
                key: state.pageKey,
              ),
            ),
          ),
          GoRoute(
            path: '/pCreateCelebration',
            name: AppRoute.pCreateCelebration.name,
            pageBuilder: (context, state) => MaterialPage(
              child: CreateCelebrationPersonal(
                key: state.pageKey,
              ),
            ),
          ),
          GoRoute(
            path: '/pTimeline',
            name: AppRoute.pTimeline.name,
            pageBuilder: (context, state) => MaterialPage(
              child: TimelinePersonal(
                key: state.pageKey,
              ),
            ),
            routes: [
              GoRoute(
                parentNavigatorKey: rootNavigation,
                path: '/pProfile',
                name: AppRoute.pProfile.name,
                pageBuilder: (context, state) {
                  return NoTransitionPage(
                    child: ProfileViewPersonal(
                      key: state.pageKey,
                    ),
                  );
                },
                routes: [
                  GoRoute(
                    parentNavigatorKey: rootNavigation,
                    path: '/pSettings',
                    name: AppRoute.pSettings.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: SettingsView(
                          key: state.pageKey,
                        ),
                      );
                    },
                    routes: [
                      GoRoute(
                        parentNavigatorKey: rootNavigation,
                        path: '/pAreasOfInterest',
                        name: AppRoute.pAreasOfInterest.name,
                        pageBuilder: (context, state) {
                          return NoTransitionPage(
                            child: AreaOfInterest(
                              key: state.pageKey,
                            ),
                          );
                        },
                      ),
                      GoRoute(
                        parentNavigatorKey: rootNavigation,
                        path: '/pNotificationSettings',
                        name: AppRoute.pNotificationSettings.name,
                        pageBuilder: (context, state) {
                          return NoTransitionPage(
                            child: NotificationsSettings(
                              key: state.pageKey,
                            ),
                          );
                        },
                      ),
                      GoRoute(
                        parentNavigatorKey: rootNavigation,
                        path: '/pOthersCelebrationsPost',
                        name: AppRoute.pOthersCelebrationsPost.name,
                        pageBuilder: (context, state) {
                          return NoTransitionPage(
                            child: OthersCelebrationPost(
                              key: state.pageKey,
                            ),
                          );
                        },
                      ),
                      GoRoute(
                        parentNavigatorKey: rootNavigation,
                        path: '/pChangePassword',
                        name: AppRoute.pChangePassword.name,
                        pageBuilder: (context, state) {
                          return NoTransitionPage(
                            child: ChangePassword(
                              key: state.pageKey,
                            ),
                          );
                        },
                      ),
                      GoRoute(
                        parentNavigatorKey: rootNavigation,
                        path: '/pPrivacyAndSafety',
                        name: AppRoute.pPrivacyAndSafety.name,
                        pageBuilder: (context, state) {
                          return NoTransitionPage(
                            child: PrivacyAndSafety(
                              key: state.pageKey,
                            ),
                          );
                        },
                      ),
                    ],
                  ),
                ],
              ),
            ],
          ),
        ],
      ),
      ShellRoute(
        navigatorKey: businessShellNavigation,
        parentNavigatorKey: rootNavigation,
        restorationScopeId: 'app',
        builder: (context, state, child) =>
            MainEntryBusinessView(key: state.pageKey, child: child),
        routes: [
          GoRoute(
            path: '/bCelebration',
            name: AppRoute.bCelebration.name,
            pageBuilder: (context, state) => MaterialPage(
              child: CelebrationBusiness(
                key: state.pageKey,
              ),
            ),
          ),
          GoRoute(
            path: '/bCreateCelebration',
            name: AppRoute.bCreateCelebration.name,
            pageBuilder: (context, state) => MaterialPage(
              child: CreateCelebrationBusiness(
                key: state.pageKey,
              ),
            ),
          ),
          GoRoute(
            path: '/bTimeline',
            name: AppRoute.bTimeline.name,
            pageBuilder: (context, state) => MaterialPage(
              child: TimelineBusiness(
                key: state.pageKey,
              ),
            ),
            routes: [
              GoRoute(
                parentNavigatorKey: rootNavigation,
                path: '/bProfile',
                name: AppRoute.bProfile.name,
                pageBuilder: (context, state) {
                  return NoTransitionPage(
                    child: ProfileViewBusiness(
                      key: state.pageKey,
                    ),
                  );
                },
                routes: [
                  GoRoute(
                    parentNavigatorKey: rootNavigation,
                    path: '/bSettings',
                    name: AppRoute.bSettings.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: SettingsView(
                          key: state.pageKey,
                        ),
                      );
                    },
                  ),
                ],
              ),
            ],
          ),
        ],
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
  verificationSuccess,
  pCreateCelebration,
  pCelebration,
  pTimeline,
  pProfile,
  pSettings,
  pAreasOfInterest,
  pChangePassword,
  pNotificationSettings,
  pOthersCelebrationsPost,
  pPrivacyAndSafety,
  bCreateCelebration,
  bCelebration,
  bTimeline,
  bProfile,
  bSettings,
}
