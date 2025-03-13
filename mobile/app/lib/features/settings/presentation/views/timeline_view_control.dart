import 'package:celebut/core/core.dart';
import 'package:celebut/features/settings/presentation/widgets/timeline_control_friend_list.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class TimelineViewControl extends StatefulWidget {
  const TimelineViewControl({super.key});

  @override
  State<TimelineViewControl> createState() => _TimelineViewControlState();
}

class _TimelineViewControlState extends State<TimelineViewControl> {
  final TextEditingController searchCtr = TextEditingController();
  String? selectedOption;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Column(
            children: [
              const SizedBox(
                height: 50,
              ),
              Row(
                children: [
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
                    'Select who can view your Timeline',
                    style: context.textTheme.titleSmall,
                  ),
                ],
              ),
              const Space(20),
              BorderAppSearchBar(
                searchCtr: searchCtr,
                labelText: 'Search Friends',
              ),
              const Space(30),
              Expanded(
                child: ListView(
                  shrinkWrap: true,
                  children: [
                    ...List.generate(
                      10,
                      (index) {
                        return TimelineControlFriendsList(
                          option: 'Daniel Bello',
                          selectedOption: selectedOption,
                          onSelected: (String value) {
                            selectedOption = value;
                            setState(() {});
                          },
                        );
                      },
                    ),
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
