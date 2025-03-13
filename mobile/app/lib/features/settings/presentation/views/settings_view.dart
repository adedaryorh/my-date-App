import 'package:celebut/core/core.dart';
import 'package:celebut/features/settings/presentation/widgets/settings_item.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class SettingsView extends StatefulWidget {
  const SettingsView({super.key});

  @override
  State<SettingsView> createState() => _SettingsViewState();
}

class _SettingsViewState extends State<SettingsView> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Settings',
          style: context.textTheme.titleMedium,
        ),
        leading: IconButton(
          onPressed: () {
            context.pop();
          },
          icon: const Icon(
            Icons.arrow_back_ios,
            size: 15,
          ),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            children: [
              const Space(30),
              ListView.builder(
                shrinkWrap: true,
                physics: const AlwaysScrollableScrollPhysics(),
                itemCount: AppConstants.settingOptions.length,
                itemBuilder: (context, index) {
                  final title = AppConstants.settingOptions[index];
                  return InkWell(
                    onTap: () {
                      // (Todo) add navigation for business apps
                      switch (title) {
                        case 'Interests':
                          context.pushNamed(AppRoute.pAreasOfInterest.name);
                        case 'Notifications':
                          context
                              .pushNamed(AppRoute.pNotificationSettings.name);
                        case 'Language':
                          context.showLanguageOptionsSheet();
                        case 'Change Password':
                          context.pushNamed(AppRoute.pChangePassword.name);
                        case 'Request for Verification':
                          break;
                        case 'Linked Devices':
                          break;
                        case 'Celebration Post by Others for me':
                          context
                              .pushNamed(AppRoute.pOthersCelebrationsPost.name);
                        case 'Surprise Event by Others for me':
                          break;
                        case 'Privacy and safety':
                          context.pushNamed(AppRoute.pPrivacyAndSafety.name);
                        case 'Log out':
                          context.showLogOutDialog();
                        default:
                      }
                    },
                    child: SettingItem(
                      title: title,
                    ),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}
