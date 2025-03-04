import 'package:celebut/apps/shared/settings/presentation/widgets/interest_options.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class AreaOfInterest extends StatefulWidget {
  const AreaOfInterest({super.key});

  @override
  State<AreaOfInterest> createState() => _AreaOfInterestState();
}

class _AreaOfInterestState extends State<AreaOfInterest> {
  final List<String> selectedInterest = [];
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(
            horizontal: 20,
          ),
          child: ListView(
            children: [
              const SizedBox(
                height: 50,
              ),
              Align(
                alignment: Alignment.centerLeft,
                child: InkWell(
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
              ),
              const SizedBox(
                height: 20,
              ),
              Text(
                'Choose Your Area of \nInterest',
                style: context.textTheme.headlineMedium,
              ),
              const Text(
                'You can customize your feed by following topics or '
                'people that interest you the most',
              ),
              const SizedBox(
                height: 30,
              ),
              Wrap(
                spacing: 5,
                runSpacing: 10,
                children: [
                  ...AppConstants.interestOptions.map((e) {
                    final isSelected = selectedInterest.contains(e);
                    return InterestOptions(
                      title: e,
                      bgColor: isSelected
                          ? context.colorScheme.primary
                          : Colors.white,
                      textColor: isSelected ? Colors.white : Colors.black,
                      onTapped: () {
                        setState(() {
                          if (isSelected) {
                            selectedInterest.remove(e);
                          } else {
                            selectedInterest.add(e);
                          }
                        });
                      },
                    );
                  }),
                ],
              ),
            ],
          ),
        ),
      ),
      bottomNavigationBar: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: MainButton(
              loading: false,
              text: 'Save',
              pressed: () {},
            ),
          ),
          const Space(30),
        ],
      ),
    );
  }
}
