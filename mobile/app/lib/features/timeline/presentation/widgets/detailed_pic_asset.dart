import 'package:flutter/material.dart';
import 'package:smooth_page_indicator/smooth_page_indicator.dart';

class DetailedPicAsset extends StatelessWidget {
  const DetailedPicAsset({
    required this.pageController,
    required this.assetPath,
    super.key,
  });

  final String assetPath;
  final PageController pageController;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Container(
          height: MediaQuery.of(context).size.height * 0.82,
          width: double.maxFinite,
          decoration: BoxDecoration(
            borderRadius: const BorderRadius.all(
              Radius.circular(20),
            ),
            color: Colors.black26,
            image: DecorationImage(
              fit: BoxFit.cover,
              image: NetworkImage(assetPath),
            ),
          ),
        ),
        Positioned(
          top: 0,
          bottom: MediaQuery.of(context).size.height * 0.82 - 200,
          left: 20,
          child: SmoothPageIndicator(
            controller: pageController,
            count: 4,
            effect: const ExpandingDotsEffect(
              dotHeight: 8,
              dotWidth: 8,
              dotColor: Colors.white70,
              activeDotColor: Colors.white,
            ),
          ),
        ),
      ],
    );
  }
}
