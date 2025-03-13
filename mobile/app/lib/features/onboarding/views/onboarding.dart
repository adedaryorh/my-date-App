import 'package:celebut/core/core.dart';
import 'package:celebut/features/onboarding/widgets/onboarding_widget.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class OnboardingView extends StatefulWidget {
  const OnboardingView({super.key});

  @override
  State<OnboardingView> createState() => _OnboardingViewState();
}

class _OnboardingViewState extends State<OnboardingView> {
  PageController pageController = PageController();
  int page = 0;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppConstants.onboardColors[page],
      body: SafeArea(
        child: ListView(
          children: [
            SizedBox(
              width: double.maxFinite,
              height: context.screenSize.height * 0.75,
              child: PageView(
                onPageChanged: (value) {
                  page = value;
                  setState(() {});
                },
                controller: pageController,
                children: getOnboardPages(pageController),
              ),
            ),
          ],
        ),
      ),
      bottomNavigationBar: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: ElevatedButton(
              onPressed: () => context.goNamed(AppRoute.accountType.name),
              style: ElevatedButton.styleFrom(backgroundColor: Colors.white),
              child: Text(
                'Next',
                style:
                    context.textTheme.bodyMedium?.copyWith(color: Colors.black),
              ),
            ),
          ),
          const Space(50),
        ],
      ),
    );
  }
}
