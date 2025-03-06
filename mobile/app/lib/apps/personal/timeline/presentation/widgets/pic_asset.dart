import 'package:flutter/material.dart';
import 'package:smooth_page_indicator/smooth_page_indicator.dart';

class PicAsset extends StatelessWidget {
  const PicAsset({
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
          margin: const EdgeInsets.symmetric(
            horizontal: 10,
            vertical: 10,
          ),
          height: 350,
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
          bottom: 300,
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

List<PicAsset> getPictures(PageController controller) {
  return [
    PicAsset(
      assetPath:
          'https://images.unsplash.com/photo-1500622944204-b135684e99fd?q=80&w=2922&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
    PicAsset(
      assetPath:
          'https://plus.unsplash.com/premium_photo-1673643405538-de0f82933fcb?q=80&w=2942&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
    PicAsset(
      assetPath:
          'https://images.unsplash.com/reserve/bOvf94dPRxWu0u3QsPjF_tree.jpg?q=80&w=2952&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
    PicAsset(
      assetPath:
          'https://plus.unsplash.com/premium_photo-1673643405538-de0f82933fcb?q=80&w=2942&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      pageController: controller,
    ),
  ];
}
