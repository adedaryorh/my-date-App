import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ControlContent extends StatefulWidget {
  const ControlContent({super.key});

  @override
  State<ControlContent> createState() => _ControlContentState();
}

class _ControlContentState extends State<ControlContent> {
  String? selectedOption;
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
              const SizedBox(
                height: 20,
              ),
              Text(
                'Content',
                style: context.textTheme.headlineMedium,
              ),
              const Text('Chose who can see things on your timeline'),
              const Space(50),
              for (int index = 0;
                  index < AppConstants.whoCanSeeContent.length;
                  index++)
                CustomRadioButton(
                  suffixWidget: index == 3
                      ? InkWell(
                          onTap: () => context
                              .pushNamed(AppRoute.timelineViewControl.name),
                          child: Text(
                            'Choose friends',
                            style: context.textTheme.bodySmall
                                ?.copyWith(color: context.colorScheme.primary),
                          ),
                        )
                      : null,
                  option: AppConstants.whoCanSeeContent[index],
                  selectedOption: selectedOption,
                  onSelected: (String value) {
                    selectedOption = value;
                    setState(() {});
                  },
                ),
            ],
          ),
        ),
      ),
    );
  }
}
