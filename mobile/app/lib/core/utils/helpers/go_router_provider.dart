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
                path: 'pMediaDetail/:mediaType',
                name: AppRoute.pMediaDetail.name,
                pageBuilder: (context, state) {
                  final mediaType = state.pathParameters['mediaType'];
                  return NoTransitionPage(
                    child: DetailMediaView(
                      mediaType: mediaType,
                      key: state.pageKey,
                    ),
                  );
                },
              ),
              GoRoute(
                parentNavigatorKey: rootNavigation,
                path: '/pNewPost',
                name: AppRoute.pNewPost.name,
                pageBuilder: (context, state) {
                  return NoTransitionPage(
                    child: NewPost(
                      key: state.pageKey,
                    ),
                  );
                },
              ),
              GoRoute(
                parentNavigatorKey: rootNavigation,
                path: '/pSharePost',
                name: AppRoute.pSharePost.name,
                pageBuilder: (context, state) {
                  return NoTransitionPage(
                    child: SharePost(
                      key: state.pageKey,
                    ),
                  );
                },
              ),
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
                    path: '/pFriendsList',
                    name: AppRoute.pFriendsList.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: FriendsList(
                          key: state.pageKey,
                        ),
                      );
                    },
                  ),
                  GoRoute(
                    parentNavigatorKey: rootNavigation,
                    path: '/pBusinessList',
                    name: AppRoute.pBusinessList.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: BusinessList(
                          key: state.pageKey,
                        ),
                      );
                    },
                  ),
                  GoRoute(
                    parentNavigatorKey: rootNavigation,
                    path: '/pEditProfile',
                    name: AppRoute.pEditProfile.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: EditProfile(
                          key: state.pageKey,
                        ),
                      );
                    },
                  ),
                  GoRoute(
                    parentNavigatorKey: rootNavigation,
                    path: '/pProfileDetails',
                    name: AppRoute.pProfileDetails.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: ProfileDetails(
                          key: state.pageKey,
                        ),
                      );
                    },
                  ),
                  GoRoute(
                    parentNavigatorKey: rootNavigation,
                    path: '/premiumDetail',
                    name: AppRoute.premiumDetail.name,
                    pageBuilder: (context, state) {
                      return NoTransitionPage(
                        child: PremiumDetail(
                          key: state.pageKey,
                        ),
                      );
                    },
                  ),
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
                        routes: [
                          GoRoute(
                            parentNavigatorKey: rootNavigation,
                            path: '/controlContent',
                            name: AppRoute.controlContent.name,
                            pageBuilder: (context, state) {
                              return NoTransitionPage(
                                child: ControlContent(
                                  key: state.pageKey,
                                ),
                              );
                            },
                          ),
                          GoRoute(
                            parentNavigatorKey: rootNavigation,
                            path: '/bannedWords',
                            name: AppRoute.bannedWords.name,
                            pageBuilder: (context, state) {
                              return NoTransitionPage(
                                child: BannedWords(
                                  key: state.pageKey,
                                ),
                              );
                            },
                          ),
                          GoRoute(
                            parentNavigatorKey: rootNavigation,
                            path: '/timelineViewControl',
                            name: AppRoute.timelineViewControl.name,
                            pageBuilder: (context, state) {
                              return NoTransitionPage(
                                child: TimelineViewControl(
                                  key: state.pageKey,
                                ),
                              );
                            },
                          ),
                          GoRoute(
                            parentNavigatorKey: rootNavigation,
                            path: '/blockedAccounts',
                            name: AppRoute.blockedAccounts.name,
                            pageBuilder: (context, state) {
                              return NoTransitionPage(
                                child: BlockedAccounts(
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
  pMediaDetail,
  pNewPost,
  pSharePost,
  pFriendsList,
  pEditProfile,
  pProfileDetails,
  pBusinessList,
  premiumDetail,
  controlContent,
  bannedWords,
  blockedAccounts,
  timelineViewControl,
  bCreateCelebration,
  bCelebration,
  bTimeline,
  bProfile,
  bSettings,
}
