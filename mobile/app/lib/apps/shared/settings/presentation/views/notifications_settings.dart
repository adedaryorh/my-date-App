import 'package:celebut/apps/shared/settings/presentation/widgets/notification_option_item.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class NotificationsSettings extends StatefulWidget {
  const NotificationsSettings({super.key});

  @override
  State<NotificationsSettings> createState() => _NotificationsSettingsState();
}

class _NotificationsSettingsState extends State<NotificationsSettings> {
  bool isSwitched = false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(
                height: 50,
              ),
              InkWell(
                onTap: () {
                  context.pop();
                },
                child: const SizedBox(
                  height: 20,
                  width: 20,
                  child: Icon(
                    Icons.arrow_back_ios,
                    size: 13,
                  ),
                ),
              ),
              const Space(20),
              Text(
                'Notification Settings',
                style: context.textTheme.headlineMedium,
              ),
              const Text('You can customize your feed by following '
                  'topics or people that interest you the most'),
              const Space(30),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(
                    'Enable Push Notifications',
                    style: context.textTheme.titleMedium
                        ?.copyWith(fontWeight: FontWeight.w400),
                  ),
                  Transform.scale(
                    scale: 0.7,
                    child: Switch(
                      value: isSwitched,
                      onChanged: (value) {
                        setState(() {
                          isSwitched = value;
                        });
                      },
                      inactiveThumbColor: Colors.grey,
                    ),
                  ),
                ],
              ),
              const Space(20),
              Container(
                height: 1,
                width: double.maxFinite,
                color: const Color(0xffE6E6E6),
              ),
              const Space(10),
              ListView(
                shrinkWrap: true,
                children: [
                  ...List.generate(
                    AppConstants.notificationOptions.length,
                    (index) {
                      return NotificationOptionItem(
                        title: AppConstants.notificationOptions[index],
                      );
                    },
                  ),
                  const SizedBox(
                    height: 50,
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
