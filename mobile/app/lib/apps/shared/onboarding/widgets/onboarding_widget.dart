import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:smooth_page_indicator/smooth_page_indicator.dart';

class OnboardingWidget extends StatelessWidget {
  const OnboardingWidget({
    required this.controller,
    required this.title,
    required this.subtitle,
    required this.imagePath,
    super.key,
  });

  final PageController controller;
  final String title;
  final String subtitle;
  final String imagePath;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const Spacer(),
        Image.asset(
          imagePath,
          height: 128,
          width: 128,
        ),
        const Space(10),
        Text(
          title,
          style: context.textTheme.headlineSmall,
          textAlign: TextAlign.center,
        ),
        const Space(10),
        Text(
          subtitle,
          style: context.textTheme.bodyLarge,
          textAlign: TextAlign.center,
        ),
        const Space(50),
        SmoothPageIndicator(
          controller: controller,
          count: 3,
          effect: const WormEffect(
              dotHeight: 8,
              dotWidth: 8,
              dotColor: Colors.white70,
              activeDotColor: Colors.white),
        ),
        const Spacer(),
      ],
    );
  }
}

List<OnboardingWidget> getOnboardPages(PageController controller) {
  return [
    OnboardingWidget(
      controller: controller,
      title: 'Celebrate Your \nLoved Ones',
      subtitle: 'Swipe right to celebrate with \npeople and in your area.',
      imagePath: AppAssets.onboardImg1,
    ),
    OnboardingWidget(
      controller: controller,
      title: 'Send Photos',
      subtitle: 'Have fun with your celebrant by '
          '\nsending photos and videos to \neach other.',
      imagePath: AppAssets.onboardImg2,
    ),
    OnboardingWidget(
      controller: controller,
      title: 'Get Notified',
      subtitle: 'Receive notifications when you \nget celebrated',
      imagePath: AppAssets.onboardImg3,
    ),
  ];
}
